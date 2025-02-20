package errno

const (
	// For api-gateway
	StatusSuccessCode = 0
	StatusSuccessMsg  = "ok"

	// For microservices
	SuccessCode = 10000
	SuccessMsg  = "ok"

	// Error
	ServiceErrorCode           = 10001 // 未知微服务错误
	ParamErrorCode             = 10002 // 参数错误
	AuthorizationFailedErrCode = 10003 // 鉴权失败
	UnexpectedTypeErrorCode    = 10004 // 未知类型
	NotImplementErrorCode      = 10005 // 未实装
	SensitiveWordsErrorCode    = 10006 // 敏感词

	// Admin
	AdminExistedErrorCode         = 1001
	AdminNotExistedErrorCode      = 1002
	UserNameOrPasswordErrorCode   = 1003
	PasswordNotCorrectErrorCode   = 1004
	UserNameIsNotEnabledErrorCode = 1005

	// metadata
	NoMetadataInContextErrorCode       = 9001
	NoAuthorizationInMetadataErrorCode = 9002

	// jwt
	InvalidAlgorithmErrorCode = 9011
	InvalidTokenErrorCode     = 9012
)
