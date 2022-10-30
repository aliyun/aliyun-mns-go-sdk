package mns

import (
	"github.com/gogap/errors"
)

const (
	ERR_NAMESPACE = "MNS"
	ERR_TEMPSTR   = "mns response status error,code: {{.resp.Code}}, message: {{.resp.Message}}, resource: {{.resource}} request id: {{.resp.RequestId}}, host id: {{.resp.HostId}}"
)

var (
	ERR_SIGN_MESSAGE_FAILED        = errors.TN(ERR_NAMESPACE, 1, "sign message failed, {{.err}}")
	ERR_MARSHAL_MESSAGE_FAILED     = errors.TN(ERR_NAMESPACE, 2, "marshal message filed, {{.err}}")
	ERR_GENERAL_AUTH_HEADER_FAILED = errors.TN(ERR_NAMESPACE, 3, "general auth header failed, {{.err}}")

	ERR_CREATE_NEW_REQUEST_FAILED = errors.TN(ERR_NAMESPACE, 4, "create new request failed, {{.err}}")
	ERR_SEND_REQUEST_FAILED       = errors.TN(ERR_NAMESPACE, 5, "send request failed, {{.err}}")
	ERR_READ_RESPONSE_BODY_FAILED = errors.TN(ERR_NAMESPACE, 6, "read response body failed, {{.err}}")

	ERR_UNMARSHAL_ERROR_RESPONSE_FAILED = errors.TN(ERR_NAMESPACE, 7, "unmarshal error response failed, {{.err}}, ResponseBody: {{.resp}}")
	ERR_UNMARSHAL_RESPONSE_FAILED       = errors.TN(ERR_NAMESPACE, 8, "unmarshal response failed, {{.err}}")
	ERR_DECODE_BODY_FAILED              = errors.TN(ERR_NAMESPACE, 9, "decode body failed, {{.err}}, body: \"{{.body}}\"")
	ERR_GET_BODY_DECODE_ELEMENT_ERROR   = errors.TN(ERR_NAMESPACE, 10, "get body decode element error, local: {{.local}}, error: {{.err}}")

	ERR_MNS_ACCESS_DENIED                = errors.TN(ERR_NAMESPACE, 100, ERR_TEMPSTR)
	ERR_MNS_INVALID_ACCESS_KEY_ID        = errors.TN(ERR_NAMESPACE, 101, ERR_TEMPSTR)
	ERR_MNS_INTERNAL_ERROR               = errors.TN(ERR_NAMESPACE, 102, ERR_TEMPSTR)
	ERR_MNS_INVALID_AUTHORIZATION_HEADER = errors.TN(ERR_NAMESPACE, 103, ERR_TEMPSTR)
	ERR_MNS_INVALID_DATE_HEADER          = errors.TN(ERR_NAMESPACE, 104, ERR_TEMPSTR)
	ERR_MNS_INVALID_ARGUMENT             = errors.TN(ERR_NAMESPACE, 105, ERR_TEMPSTR)
	ERR_MNS_INVALID_DEGIST               = errors.TN(ERR_NAMESPACE, 106, ERR_TEMPSTR)
	ERR_MNS_INVALID_REQUEST_URL          = errors.TN(ERR_NAMESPACE, 107, ERR_TEMPSTR)
	ERR_MNS_INVALID_QUERY_STRING         = errors.TN(ERR_NAMESPACE, 108, ERR_TEMPSTR)
	ERR_MNS_MALFORMED_XML                = errors.TN(ERR_NAMESPACE, 109, ERR_TEMPSTR)
	ERR_MNS_MISSING_AUTHORIZATION_HEADER = errors.TN(ERR_NAMESPACE, 110, ERR_TEMPSTR)
	ERR_MNS_MISSING_DATE_HEADER          = errors.TN(ERR_NAMESPACE, 111, ERR_TEMPSTR)
	ERR_MNS_MISSING_VERSION_HEADER       = errors.TN(ERR_NAMESPACE, 112, ERR_TEMPSTR)
	ERR_MNS_MISSING_RECEIPT_HANDLE       = errors.TN(ERR_NAMESPACE, 113, ERR_TEMPSTR)
	ERR_MNS_MISSING_VISIBILITY_TIMEOUT   = errors.TN(ERR_NAMESPACE, 114, ERR_TEMPSTR)
	ERR_MNS_MESSAGE_NOT_EXIST            = errors.TN(ERR_NAMESPACE, 115, ERR_TEMPSTR)
	ERR_MNS_QUEUE_ALREADY_EXIST          = errors.TN(ERR_NAMESPACE, 116, ERR_TEMPSTR)
	ERR_MNS_QUEUE_DELETED_RECENTLY       = errors.TN(ERR_NAMESPACE, 117, ERR_TEMPSTR)
	ERR_MNS_INVALID_QUEUE_NAME           = errors.TN(ERR_NAMESPACE, 118, ERR_TEMPSTR)
	ERR_MNS_INVALID_VERSION_HEADER       = errors.TN(ERR_NAMESPACE, 119, ERR_TEMPSTR)
	ERR_MNS_INVALID_CONTENT_TYPE         = errors.TN(ERR_NAMESPACE, 120, ERR_TEMPSTR)
	ERR_MNS_QUEUE_NAME_LENGTH_ERROR      = errors.TN(ERR_NAMESPACE, 121, ERR_TEMPSTR)
	ERR_MNS_QUEUE_NOT_EXIST              = errors.TN(ERR_NAMESPACE, 122, ERR_TEMPSTR)
	ERR_MNS_RECEIPT_HANDLE_ERROR         = errors.TN(ERR_NAMESPACE, 123, ERR_TEMPSTR)
	ERR_MNS_SIGNATURE_DOES_NOT_MATCH     = errors.TN(ERR_NAMESPACE, 124, ERR_TEMPSTR)
	ERR_MNS_TIME_EXPIRED                 = errors.TN(ERR_NAMESPACE, 125, ERR_TEMPSTR)
	ERR_MNS_QPS_LIMIT_EXCEEDED           = errors.TN(ERR_NAMESPACE, 134, ERR_TEMPSTR)
	ERR_MNS_UNKNOWN_CODE                 = errors.TN(ERR_NAMESPACE, 135, ERR_TEMPSTR)

	ERR_MNS_TOPIC_NAME_LENGTH_ERROR       = errors.TN(ERR_NAMESPACE, 200, ERR_TEMPSTR)
	ERR_MNS_SUBSRIPTION_NAME_LENGTH_ERROR = errors.TN(ERR_NAMESPACE, 201, ERR_TEMPSTR)
	ERR_MNS_TOPIC_NOT_EXIST               = errors.TN(ERR_NAMESPACE, 202, ERR_TEMPSTR)
	ERR_MNS_TOPIC_ALREADY_EXIST           = errors.TN(ERR_NAMESPACE, 203, ERR_TEMPSTR)
	ERR_MNS_INVALID_TOPIC_NAME            = errors.TN(ERR_NAMESPACE, 204, ERR_TEMPSTR)
	ERR_MNS_INVALID_SUBSCRIPTION_NAME     = errors.TN(ERR_NAMESPACE, 205, ERR_TEMPSTR)
	ERR_MNS_SUBSCRIPTION_ALREADY_EXIST    = errors.TN(ERR_NAMESPACE, 206, ERR_TEMPSTR)
	ERR_MNS_INVALID_ENDPOINT              = errors.TN(ERR_NAMESPACE, 207, ERR_TEMPSTR)
	ERR_MNS_SUBSCRIBER_NOT_EXIST          = errors.TN(ERR_NAMESPACE, 211, ERR_TEMPSTR)

	ERR_MNS_TOPIC_NAME_IS_TOO_LONG                        = errors.TN(ERR_NAMESPACE, 208, "topic name is too long, the max length is 256")
	ERR_MNS_TOPIC_ALREADY_EXIST_AND_HAVE_SAME_ATTR        = errors.TN(ERR_NAMESPACE, 209, "mns topic already exist, and the attribute is the same, topic name: {{.name}}")
	ERR_MNS_SUBSCRIPTION_ALREADY_EXIST_AND_HAVE_SAME_ATTR = errors.TN(ERR_NAMESPACE, 210, "mns subscription already exist, and the attribute is the same, subscription name: {{.name}}")

	ERR_MNS_QUEUE_NAME_IS_TOO_LONG                 = errors.TN(ERR_NAMESPACE, 126, "queue name is too long, the max length is 256")
	ERR_MNS_DELAY_SECONDS_RANGE_ERROR              = errors.TN(ERR_NAMESPACE, 127, "queue delay seconds is not in range of (0~60480)")
	ERR_MNS_MAX_MESSAGE_SIZE_RANGE_ERROR           = errors.TN(ERR_NAMESPACE, 128, "max message size is not in range of (1024~65536)")
	ERR_MNS_MSG_RETENTION_PERIOD_RANGE_ERROR       = errors.TN(ERR_NAMESPACE, 129, "message retention period is not in range of (60~129600)")
	ERR_MNS_MSG_VISIBILITY_TIMEOUT_RANGE_ERROR     = errors.TN(ERR_NAMESPACE, 130, "message visibility timeout is not in range of (1~43200)")
	ERR_MNS_MSG_POOLLING_WAIT_SECONDS_RANGE_ERROR  = errors.TN(ERR_NAMESPACE, 131, "message poolling wait seconds is not in range of (0~30)")
	ERR_MNS_RET_NUMBER_RANGE_ERROR                 = errors.TN(ERR_NAMESPACE, 132, "list param of ret number is not in range of (1~1000)")
	ERR_MNS_QUEUE_ALREADY_EXIST_AND_HAVE_SAME_ATTR = errors.TN(ERR_NAMESPACE, 133, "mns queue already exist, and the attribute is the same, queue name: {{.name}}")
	ERR_MNS_BATCH_OP_FAIL                          = errors.TN(ERR_NAMESPACE, 136, "mns queue batch operation fail")
)
