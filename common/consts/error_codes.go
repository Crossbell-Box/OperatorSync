package consts

// Code rule:
// 1 xx xx
// |  |  |
// |  |  Error No
// |  Group No
// Generation No (API v1 / v2 /...)

const (
	ERROR_CODE_UNSUPPORTED_PLATFORM           = 10101 // Submit request errors
	ERROR_CODE_INVALID_FORMAT                 = 10102
	ERROR_CODE_HTTP_REQUEST_FAILED            = 10201 // Request errors
	ERROR_CODE_FAILED_TO_PARSE_FEEDS          = 10202
	ERROR_CODE_FAILED_TO_FIND_NECESSARY_FIELD = 10203
	ERROR_CODE_FAILED_TO_PARSE_JSON           = 10301 // System internal errors
	ERROR_CODE_FAILED_TO_UPLOAD               = 10401 // External system errors (like rate limit)

)
