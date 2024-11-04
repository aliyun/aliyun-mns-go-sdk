package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

func main() {
	headers := map[string]string{
		"Connection":              "keep-alive",
		"x-mns-message-id":        "AC10020B0019681A95159E591D10****",
		"Content-Type":            "text/plain;charset=utf-8",
		"Content-md5":             "5B4682CCA07FFB080FFE0A9D9821****",
		"x-mns-topic-owner":       "****",
		"x-mns-topic-name":        "HTTP-test",
		"x-mns-subscriber":        "****",
		"x-mns-subscription-name": "sub-test-header",
		"x-mns-publish-time":      "1730368640272",
		"x-mns-request-id":        "672354803035301900003***",
		"Date":                    "Thu, 31 Oct 2024 09:57:20 GMT",
		"x-mns-version":           "2015-06-06",
		"User-Agent":              "Aliyun Notification Service Agent",
		"x-mns-signing-cert-url":  "aHR0cHM6Ly9tbnN0ZXN0Lm9zcy1jbi1oYW5nemhvdS5hbGl5dW5****",
		"Authorization":           "pg5Prc+ADujqjHbK1XKMK+o+aZjtkAntpR19s2B0T1k1deilZ5UgUFoIsKmLbgirN+1m2srdh****",
		"Host":                    "****",
		"Content-length":          "40",
		"Accept":                  "*/*",
	}

	method := "POST"
	path := "/api/test"
	if authenticateWithHeaderMap(method, path, toLowercaseKeys(headers)) {
		fmt.Println("Signature verification succeeded")
	} else {
		fmt.Println("Signature verification failed")
	}
}

func authenticateWithResponse(method, path string, resp *http.Response) bool {
	headersMap := headersToMap(resp)
	return authenticateWithHeaderMap(method, path, headersMap)
}

func authenticateWithHeaderMap(method, path string, headers map[string]string) bool {
	// Get string to sign
	var serviceHeaders []string
	for k, v := range headers {
		if strings.HasPrefix(strings.ToLower(k), "x-mns-") {
			serviceHeaders = append(serviceHeaders, fmt.Sprintf("%s:%s", k, v))
		}
	}

	sort.Strings(serviceHeaders)
	serviceStr := strings.Join(serviceHeaders, "\n")
	var signHeaderList []string
	for _, key := range []string{"content-md5", "content-type", "date"} {
		if val, ok := headers[key]; ok {
			signHeaderList = append(signHeaderList, val)
		} else {
			signHeaderList = append(signHeaderList, "")
		}
	}

	str2sign := fmt.Sprintf("%s\n%s\n%s\n%s", method, strings.Join(signHeaderList, "\n"), serviceStr, path)
	fmt.Println("String to sign:", str2sign)

	// 获取证书的URL
	certURLBase64, ok := headers["x-mns-signing-cert-url"]
	if !ok {
		fmt.Println("x-mns-signing-cert-url Header not found")
		return false
	}

	certURLBytes, err := base64.StdEncoding.DecodeString(certURLBase64)
	if err != nil {
		fmt.Println("Failed to decode base64 cert URL:", err)
		return false
	}

	certURL := string(certURLBytes)
	fmt.Println("x-mns-signing-cert-url:\t", certURL)

	// 根据URL获取证书，并从证书中获取公钥
	resp, err := http.Get(certURL)
	if err != nil {
		fmt.Println("Failed to fetch certificate:", err)
		return false
	}
	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()

	//goland:noinspection GoDeprecation
	certData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read certificate:", err)
		return false
	}

	block, _ := pem.Decode(certData)
	if block == nil || block.Type != "CERTIFICATE" {
		fmt.Println("Failed to decode PEM block containing the certificate")
		return false
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println("Failed to parse certificate:", err)
		return false
	}

	pubKey, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		fmt.Println("Failed to cast public key to RSA public key")
		return false
	}

	// 对Authorization字段做Base64解码
	signatureBase64, ok := headers["authorization"]
	if !ok {
		fmt.Println("Authorization Header not found")
		return false
	}

	signature, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		fmt.Println("Failed to decode base64 signature:", err)
		return false
	}

	// 认证
	hash := sha1.New()
	hash.Write([]byte(str2sign))
	digest := hash.Sum(nil)
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA1, digest, signature)
	if err != nil {
		fmt.Println("Signature verification failed:", err)
		return false
	}

	return true
}

func headersToMap(resp *http.Response) map[string]string {
	headersMap := make(map[string]string)
	for key, values := range resp.Header {
		// 连接多个值为一个字符串，使用逗号分隔
		// map 键值全小写
		lowercaseKey := strings.ToLower(key)
		headersMap[lowercaseKey] = strings.Join(values, ",")
	}
	return headersMap
}

// HTTP header 不区分大小写，将 map 的 keys 转换为全小写
func toLowercaseKeys(input map[string]string) map[string]string {
	output := make(map[string]string)
	for k, v := range input {
		lowercaseKey := strings.ToLower(k)
		output[lowercaseKey] = v
	}
	return output
}
