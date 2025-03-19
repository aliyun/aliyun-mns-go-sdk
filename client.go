package ali_mns

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"net/http"
	neturl "net/url"
	"os"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/aliyun/credentials-go/credentials"
	"github.com/gogap/errors"
	"github.com/valyala/fasthttp"
)

const (
	DefaultQueueQPSLimit int32 = 2000
	DefaultTopicQPSLimit int32 = 2000
	DefaultDNSTTL        int32 = 10
)

const (
	GlobalProxy = "MNS_GLOBAL_PROXY"
)

const (
	version = "2015-06-06"
)

const (
	DefaultTimeout         int64 = 35
	DefaultMaxConnsPerHost int   = 512
)

const (
	AliyunAkEnvKey = "ALIBABA_CLOUD_ACCESS_KEY_ID"
	AliyunSkEnvKey = "ALIBABA_CLOUD_ACCESS_KEY_SECRET"
)

type Method string

var (
	errMapping map[string]errors.ErrCodeTemplate
)

func init() {
	initMNSErrors()
}

const (
	GET    Method = "GET"
	PUT           = "PUT"
	POST          = "POST"
	DELETE        = "DELETE"
)

type MNSClient interface {
	Send(method Method, headers map[string]string, message interface{}, resource string) (*fasthttp.Response, error)
	SetProxy(url string)
	SetTransport(transport fasthttp.RoundTripper)
	getAccountId() (accountId string)
	getRegion() (region string)
}

type aliMNSClient struct {
	Timeout         int64
	MaxConnsPerHost int
	url             *neturl.URL
	credential      credentials.Credential
	accessKeyId     string
	client          *fasthttp.Client
	proxyURL        string
	accountId       string
	region          string
	clientLocker    sync.Mutex
}

type AliMNSClientConfig struct {
	EndPoint        string
	AccessKeyId     string
	AccessKeySecret string
	Token           string
	Credential      credentials.Credential
	TimeoutSecond   int64
	MaxConnsPerHost int
}

// NewClient Follow the Alibaba Cloud standards and set the AK (Access Key) and SK (Secret Key) in the environment variables.
// For more details, see: https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems
func NewClient(endpoint string) MNSClient {
	return NewClientWithToken(endpoint, "")
}

// NewClientWithToken Follow the Alibaba Cloud standards and set the AK (Access Key) and SK (Secret Key) in the environment variables.
// For more details, see: https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems
func NewClientWithToken(endpoint, token string) MNSClient {
	return NewAliMNSClientWithConfig(AliMNSClientConfig{
		EndPoint:        endpoint,
		AccessKeyId:     os.Getenv(AliyunAkEnvKey),
		AccessKeySecret: os.Getenv(AliyunSkEnvKey),
		Token:           token,
		TimeoutSecond:   DefaultTimeout,
		MaxConnsPerHost: DefaultMaxConnsPerHost,
	})
}

// Deprecated: Use NewClient instead.
func NewAliMNSClient(inputUrl, accessKeyId, accessKeySecret string) MNSClient {
	return NewAliMNSClientWithConfig(AliMNSClientConfig{
		EndPoint:        inputUrl,
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Token:           "",
		TimeoutSecond:   DefaultTimeout,
		MaxConnsPerHost: DefaultMaxConnsPerHost,
	})
}

// Deprecated: Use NewClientWithToken instead.
func NewAliMNSClientWithToken(inputUrl, accessKeyId, accessKeySecret, token string) MNSClient {
	return NewAliMNSClientWithConfig(AliMNSClientConfig{
		EndPoint:        inputUrl,
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Token:           token,
		TimeoutSecond:   DefaultTimeout,
		MaxConnsPerHost: DefaultMaxConnsPerHost,
	})
}

func NewAliMNSClientWithConfig(clientConfig AliMNSClientConfig) MNSClient {
	if clientConfig.EndPoint == "" {
		panic("ali-mns: message queue url is empty")
	}

	cli := new(aliMNSClient)
	cli.Timeout = clientConfig.TimeoutSecond
	if clientConfig.Credential != nil {
		cli.credential = clientConfig.Credential
	} else if clientConfig.Token != "" {
		config := new(credentials.Config).
			SetType("sts").
			SetAccessKeyId(clientConfig.AccessKeyId).
			SetAccessKeySecret(clientConfig.AccessKeySecret).
			SetSecurityToken(clientConfig.Token)
		var err error
		cli.credential, err = credentials.NewCredential(config)
		if err != nil {
			panic(err)
		}
	} else {
		config := new(credentials.Config).
			SetType("access_key").
			SetAccessKeyId(clientConfig.AccessKeyId).
			SetAccessKeySecret(clientConfig.AccessKeySecret)
		var err error
		cli.credential, err = credentials.NewCredential(config)
		if err != nil {
			panic(err)
		}
	}

	if clientConfig.MaxConnsPerHost != 0 {
		cli.MaxConnsPerHost = clientConfig.MaxConnsPerHost
	} else {
		cli.MaxConnsPerHost = DefaultMaxConnsPerHost
	}

	var err error
	if cli.url, err = neturl.Parse(clientConfig.EndPoint); err != nil {
		panic("err parse url")
	}

	// 1. parse region and accountId
	pieces := strings.Split(clientConfig.EndPoint, ".")
	if len(pieces) != 5 {
		panic("ali-mns: message queue url is invalid")
	}

	accountIdSlice := strings.Split(pieces[0], "/")
	cli.accountId = accountIdSlice[len(accountIdSlice)-1]
	re := regexp.MustCompile("-(internal|control)")
	regionSlice := re.Split(pieces[2], -1)
	cli.region = regionSlice[0]
	if globalUrl := os.Getenv(GlobalProxy); globalUrl != "" {
		cli.proxyURL = globalUrl
	}

	// 2. now init http client
	cli.initFastHttpClient()
	//change to dial dual stack to support both ipv4 and ipv6
	cli.client.DialDualStack = true
	return cli
}

func (p aliMNSClient) getAccountId() (accountId string) {
	return p.accountId
}

func (p aliMNSClient) getRegion() (region string) {
	return p.region
}

func (p *aliMNSClient) SetProxy(url string) {
	if url == p.proxyURL {
		return
	}

	p.proxyURL = url
}

func (p *aliMNSClient) initFastHttpClient() {
	p.clientLocker.Lock()
	defer p.clientLocker.Unlock()
	timeoutInt := DefaultTimeout
	if p.Timeout > 0 {
		timeoutInt = p.Timeout
	}

	timeout := time.Second * time.Duration(timeoutInt)
	p.client = &fasthttp.Client{ReadTimeout: timeout, WriteTimeout: timeout, MaxConnsPerHost: p.MaxConnsPerHost, Name: getDefaultUserAgent()}
}

func (p *aliMNSClient) SetTransport(transport fasthttp.RoundTripper) {
	p.client.ConfigureClient = func(hc *fasthttp.HostClient) error {
		hc.Transport = transport
		return nil
	}
}

func (p *aliMNSClient) proxy() (*neturl.URL, error) {
	if p.proxyURL != "" {
		return neturl.Parse(p.proxyURL)
	}
	return nil, nil
}

func (p *aliMNSClient) Send(method Method, headers map[string]string, message interface{}, resource string) (*fasthttp.Response, error) {
	var xmlContent []byte
	var err error

	if message == nil {
		xmlContent = []byte{}
	} else {
		switch m := message.(type) {
		case []byte:
			{
				xmlContent = m
			}
		default:
			messageSendRequest, ok := message.(MessageSendRequest)
			if ok && messageSendRequest.Priority == 0 {
				messageSendRequest.Priority = 8
			}
			if bXml, e := xml.Marshal(messageSendRequest); e != nil {
				err = ERR_MARSHAL_MESSAGE_FAILED.New(errors.Params{"err": e})
				return nil, err
			} else {
				xmlContent = bXml
			}
		}
	}

	xmlMD5 := md5.Sum(xmlContent)
	strMd5 := fmt.Sprintf("%x", xmlMD5)

	if headers == nil {
		headers = make(map[string]string)
	}

	headers[MQ_VERSION] = version
	headers[CONTENT_TYPE] = "application/xml"
	headers[CONTENT_MD5] = base64.StdEncoding.EncodeToString([]byte(strMd5))
	headers[DATE] = time.Now().UTC().Format(http.TimeFormat)

	credential, err := p.credential.GetCredential()
	if err != nil {
		return nil, err
	}
	if credential.SecurityToken != nil && *credential.SecurityToken != "" {
		headers[SECURITY_TOKEN] = *credential.SecurityToken
	}

	signature, err := getSignature(method, headers, fmt.Sprintf("/%s", resource), *credential.AccessKeySecret)
	if err != nil {
		return nil, ERR_GENERAL_AUTH_HEADER_FAILED.New(errors.Params{"err": err})
	}
	headers[AUTHORIZATION] = fmt.Sprintf("MNS %s:%s", *credential.AccessKeyId, signature)

	var buffer bytes.Buffer
	buffer.WriteString(p.url.String())
	buffer.WriteString("/")
	buffer.WriteString(resource)

	url := buffer.String()

	req := fasthttp.AcquireRequest()

	req.SetRequestURI(url)
	req.Header.SetMethod(string(method))
	req.SetBody(xmlContent)

	for header, value := range headers {
		req.Header.Set(header, value)
	}

	resp := fasthttp.AcquireResponse()

	if err = p.client.Do(req, resp); err != nil {
		err = ERR_SEND_REQUEST_FAILED.New(errors.Params{"err": err})
		return nil, err
	}

	return resp, nil
}

func initMNSErrors() {
	errMapping = map[string]errors.ErrCodeTemplate{
		"AccessDenied":                ERR_MNS_ACCESS_DENIED,
		"InvalidAccessKeyId":          ERR_MNS_INVALID_ACCESS_KEY_ID,
		"InternalError":               ERR_MNS_INTERNAL_ERROR,
		"InvalidAuthorizationHeader":  ERR_MNS_INVALID_AUTHORIZATION_HEADER,
		"InvalidDateHeader":           ERR_MNS_INVALID_DATE_HEADER,
		"InvalidArgument":             ERR_MNS_INVALID_ARGUMENT,
		"InvalidDegist":               ERR_MNS_INVALID_DEGIST,
		"InvalidRequestURL":           ERR_MNS_INVALID_REQUEST_URL,
		"InvalidQueryString":          ERR_MNS_INVALID_QUERY_STRING,
		"MalformedXML":                ERR_MNS_MALFORMED_XML,
		"MissingAuthorizationHeader":  ERR_MNS_MISSING_AUTHORIZATION_HEADER,
		"MissingDateHeader":           ERR_MNS_MISSING_DATE_HEADER,
		"MissingVersionHeader":        ERR_MNS_MISSING_VERSION_HEADER,
		"MissingReceiptHandle":        ERR_MNS_MISSING_RECEIPT_HANDLE,
		"MissingVisibilityTimeout":    ERR_MNS_MISSING_VISIBILITY_TIMEOUT,
		"MessageNotExist":             ERR_MNS_MESSAGE_NOT_EXIST,
		"QueueAlreadyExist":           ERR_MNS_QUEUE_ALREADY_EXIST,
		"QueueDeletedRecently":        ERR_MNS_QUEUE_DELETED_RECENTLY,
		"InvalidQueueName":            ERR_MNS_INVALID_QUEUE_NAME,
		"QueueNameLengthError":        ERR_MNS_QUEUE_NAME_LENGTH_ERROR,
		"QueueNotExist":               ERR_MNS_QUEUE_NOT_EXIST,
		"ReceiptHandleError":          ERR_MNS_RECEIPT_HANDLE_ERROR,
		"SignatureDoesNotMatch":       ERR_MNS_SIGNATURE_DOES_NOT_MATCH,
		"TimeExpired":                 ERR_MNS_TIME_EXPIRED,
		"QpsLimitExceeded":            ERR_MNS_QPS_LIMIT_EXCEEDED,
		"TopicAlreadyExist":           ERR_MNS_TOPIC_ALREADY_EXIST,
		"TopicNameLengthError":        ERR_MNS_TOPIC_NAME_LENGTH_ERROR,
		"TopicNotExist":               ERR_MNS_TOPIC_NOT_EXIST,
		"SubscriptionNameLengthError": ERR_MNS_SUBSRIPTION_NAME_LENGTH_ERROR,
		"TopicNameInvalid":            ERR_MNS_INVALID_TOPIC_NAME,
		"SubsriptionNameInvalid":      ERR_MNS_INVALID_SUBSCRIPTION_NAME,
		"SubscriptionAlreadyExist":    ERR_MNS_SUBSCRIPTION_ALREADY_EXIST,
		"EndpointInvalid":             ERR_MNS_INVALID_ENDPOINT,
		"SubscriberNotExist":          ERR_MNS_SUBSCRIBER_NOT_EXIST,
	}
}

func getDefaultUserAgent() string {
	goVersion := strings.TrimPrefix(runtime.Version(), "go")
	return fmt.Sprintf("%s/%s(%s/%s/%s;%s)", SdkName, Version, runtime.GOOS, "-", runtime.GOARCH, goVersion)
}

func ParseError(resp ErrorResponse, resource string) (err error) {
	if errCodeTemplate, exist := errMapping[resp.Code]; exist {
		err = errCodeTemplate.New(errors.Params{"resp": resp, "resource": resource})
	} else {
		err = ERR_MNS_UNKNOWN_CODE.New(errors.Params{"resp": resp, "resource": resource})
	}
	return
}
