package mns

import (
	"github.com/gogap/errors"
)

const (
	mns_ERR_NS = "MNS"

	mns_ERR_TEMPSTR = "mns response status error,code: {{.resp.Code}}, message: {{.resp.Message}}, resource: {{.resource}} request id: {{.resp.RequestId}}, host id: {{.resp.HostId}}"
)

var (
	ERR_SIGN_MESSAGE_FAILED        = errors.TN(mns_ERR_NS, 1, "sign message failed, {{.err}}")
	ERR_MARSHAL_MESSAGE_FAILED     = errors.TN(mns_ERR_NS, 2, "marshal message filed, {{.err}}")
	ERR_GENERAL_AUTH_HEADER_FAILED = errors.TN(mns_ERR_NS, 3, "general auth header failed, {{.err}}")

	ERR_CREATE_NEW_REQUEST_FAILED = errors.TN(mns_ERR_NS, 4, "create new request failed, {{.err}}")
	ERR_SEND_REQUEST_FAILED       = errors.TN(mns_ERR_NS, 5, "send request failed, {{.err}}")
	ERR_READ_RESPONSE_BODY_FAILED = errors.TN(mns_ERR_NS, 6, "read response body failed, {{.err}}")

	ERR_UNMARSHAL_ERROR_RESPONSE_FAILED = errors.TN(mns_ERR_NS, 7, "unmarshal error response failed, {{.err}}, ResponseBody: {{.resp}}")
	ERR_UNMARSHAL_RESPONSE_FAILED       = errors.TN(mns_ERR_NS, 8, "unmarshal response failed, {{.err}}")
	ERR_DECODE_BODY_FAILED              = errors.TN(mns_ERR_NS, 9, "decode body failed, {{.err}}, body: \"{{.body}}\"")
	ERR_GET_BODY_DECODE_ELEMENT_ERROR   = errors.TN(mns_ERR_NS, 10, "get body decode element error, local: {{.local}}, error: {{.err}}")

	ERR_MNS_ACCESS_DENIED                = errors.TN(mns_ERR_NS, 100, mns_ERR_TEMPSTR)
	ERR_MNS_INVALID_ACCESS_KEY_ID        = errors.TN(mns_ERR_NS, 101, mns_ERR_TEMPSTR)
	ERR_MNS_INTERNAL_ERROR               = errors.TN(mns_ERR_NS, 102, mns_ERR_TEMPSTR)
	ERR_MNS_INVALID_AUTHORIZATION_HEADER = errors.TN(mns_ERR_NS, 103, mns_ERR_TEMPSTR)
	ERR_MNS_INVALID_DATE_HEADER          = errors.TN(mns_ERR_NS, 104, mns_ERR_TEMPSTR)
	ERR_MNS_INVALID_ARGUMENT             = errors.TN(mns_ERR_NS, 105, mns_ERR_TEMPSTR)
	ERR_MNS_INVALID_DEGIST               = errors.TN(mns_ERR_NS, 106, mns_ERR_TEMPSTR)
	ERR_MNS_INVALID_REQUEST_URL          = errors.TN(mns_ERR_NS, 107, mns_ERR_TEMPSTR)
	ERR_MNS_INVALID_QUERY_STRING         = errors.TN(mns_ERR_NS, 108, mns_ERR_TEMPSTR)
	ERR_MNS_MALFORMED_XML                = errors.TN(mns_ERR_NS, 109, mns_ERR_TEMPSTR)
	ERR_MNS_MISSING_AUTHORIZATION_HEADER = errors.TN(mns_ERR_NS, 110, mns_ERR_TEMPSTR)
	ERR_MNS_MISSING_DATE_HEADER          = errors.TN(mns_ERR_NS, 111, mns_ERR_TEMPSTR)
	ERR_MNS_MISSING_VERSION_HEADER       = errors.TN(mns_ERR_NS, 112, mns_ERR_TEMPSTR)
	ERR_MNS_MISSING_RECEIPT_HANDLE       = errors.TN(mns_ERR_NS, 113, mns_ERR_TEMPSTR)
	ERR_MNS_MISSING_VISIBILITY_TIMEOUT   = errors.TN(mns_ERR_NS, 114, mns_ERR_TEMPSTR)
	ERR_MNS_MESSAGE_NOT_EXIST            = errors.TN(mns_ERR_NS, 115, mns_ERR_TEMPSTR)
	ERR_MNS_QUEUE_ALREADY_EXIST          = errors.TN(mns_ERR_NS, 116, mns_ERR_TEMPSTR)
	ERR_MNS_QUEUE_DELETED_RECENTLY       = errors.TN(mns_ERR_NS, 117, mns_ERR_TEMPSTR)
	ERR_MNS_INVALID_QUEUE_NAME           = errors.TN(mns_ERR_NS, 118, mns_ERR_TEMPSTR)
	ERR_MNS_INVALID_VERSION_HEADER       = errors.TN(mns_ERR_NS, 119, mns_ERR_TEMPSTR)
	ERR_MNS_INVALID_CONTENT_TYPE         = errors.TN(mns_ERR_NS, 120, mns_ERR_TEMPSTR)
	ERR_MNS_QUEUE_NAME_LENGTH_ERROR      = errors.TN(mns_ERR_NS, 121, mns_ERR_TEMPSTR)
	ERR_MNS_QUEUE_NOT_EXIST              = errors.TN(mns_ERR_NS, 122, mns_ERR_TEMPSTR)
	ERR_MNS_RECEIPT_HANDLE_ERROR         = errors.TN(mns_ERR_NS, 123, mns_ERR_TEMPSTR)
	ERR_MNS_SIGNATURE_DOES_NOT_MATCH     = errors.TN(mns_ERR_NS, 124, mns_ERR_TEMPSTR)
	ERR_MNS_TIME_EXPIRED                 = errors.TN(mns_ERR_NS, 125, mns_ERR_TEMPSTR)
	ERR_MNS_QPS_LIMIT_EXCEEDED           = errors.TN(mns_ERR_NS, 134, mns_ERR_TEMPSTR)
	ERR_MNS_UNKNOWN_CODE                 = errors.TN(mns_ERR_NS, 135, mns_ERR_TEMPSTR)

	ERR_MNS_TOPIC_NAME_LENGTH_ERROR       = errors.TN(mns_ERR_NS, 200, mns_ERR_TEMPSTR)
	ERR_MNS_SUBSRIPTION_NAME_LENGTH_ERROR = errors.TN(mns_ERR_NS, 201, mns_ERR_TEMPSTR)
	ERR_MNS_TOPIC_NOT_EXIST               = errors.TN(mns_ERR_NS, 202, mns_ERR_TEMPSTR)
	ERR_MNS_TOPIC_ALREADY_EXIST           = errors.TN(mns_ERR_NS, 203, mns_ERR_TEMPSTR)
	ERR_MNS_INVALID_TOPIC_NAME            = errors.TN(mns_ERR_NS, 204, mns_ERR_TEMPSTR)
	ERR_MNS_INVALID_SUBSCRIPTION_NAME     = errors.TN(mns_ERR_NS, 205, mns_ERR_TEMPSTR)
	ERR_MNS_SUBSCRIPTION_ALREADY_EXIST    = errors.TN(mns_ERR_NS, 206, mns_ERR_TEMPSTR)
	ERR_MNS_INVALID_ENDPOINT              = errors.TN(mns_ERR_NS, 207, mns_ERR_TEMPSTR)
	ERR_MNS_SUBSCRIBER_NOT_EXIST          = errors.TN(mns_ERR_NS, 211, mns_ERR_TEMPSTR)

	ERR_MNS_TOPIC_NAME_IS_TOO_LONG                        = errors.TN(mns_ERR_NS, 208, "topic name is too long, the max length is 256")
	ERR_MNS_TOPIC_ALREADY_EXIST_AND_HAVE_SAME_ATTR        = errors.TN(mns_ERR_NS, 209, "mns topic already exist, and the attribute is the same, topic name: {{.name}}")
	ERR_MNS_SUBSCRIPTION_ALREADY_EXIST_AND_HAVE_SAME_ATTR = errors.TN(mns_ERR_NS, 210, "mns subscription already exist, and the attribute is the same, subscription name: {{.name}}")

	ERR_MNS_QUEUE_NAME_IS_TOO_LONG                 = errors.TN(mns_ERR_NS, 126, "queue name is too long, the max length is 256")
	ERR_MNS_DELAY_SECONDS_RANGE_ERROR              = errors.TN(mns_ERR_NS, 127, "queue delay seconds is not in range of (0~60480)")
	ERR_MNS_MAX_MESSAGE_SIZE_RANGE_ERROR           = errors.TN(mns_ERR_NS, 128, "max message size is not in range of (1024~65536)")
	ERR_MNS_MSG_RETENTION_PERIOD_RANGE_ERROR       = errors.TN(mns_ERR_NS, 129, "message retention period is not in range of (60~129600)")
	ERR_MNS_MSG_VISIBILITY_TIMEOUT_RANGE_ERROR     = errors.TN(mns_ERR_NS, 130, "message visibility timeout is not in range of (1~43200)")
	ERR_MNS_MSG_POOLLING_WAIT_SECONDS_RANGE_ERROR  = errors.TN(mns_ERR_NS, 131, "message poolling wait seconds is not in range of (0~30)")
	ERR_MNS_RET_NUMBER_RANGE_ERROR                 = errors.TN(mns_ERR_NS, 132, "list param of ret number is not in range of (1~1000)")
	ERR_MNS_QUEUE_ALREADY_EXIST_AND_HAVE_SAME_ATTR = errors.TN(mns_ERR_NS, 133, "mns queue already exist, and the attribute is the same, queue name: {{.name}}")
	ERR_MNS_BATCH_OP_FAIL                          = errors.TN(mns_ERR_NS, 136, "mns queue batch operation fail")
)
