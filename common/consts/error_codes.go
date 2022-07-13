package consts

// Code rule:
// 1 xx xx
// |  |  |
// |  |  Error No
// |  Group No
// Generation No (API v1 / v2 /...)

const (
	ERROR_CODE_UNSUPPORTED_PLATFORM  = 10101 // Submit request errors
	ERROR_CODE_HTTP_REQUEST_FAILED   = 10201 // System internal errors
	ERROR_CODE_FAILED_TO_PARSE_FEEDS = 10202
	ERROR_CODE_ACCOUNT_NOT_VERIFIED  = 10301 // User setting errors
	ERROR_CODE_NOTHING_FOUND         = 10302 // Nothing-new found
)
