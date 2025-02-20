package errno

var (
	// Success
	Success = NewErrNo(SuccessCode, "Success")

	ServiceError             = NewErrNo(ServiceErrorCode, "service is unable to start successfully")
	ServiceInternalError     = NewErrNo(ServiceErrorCode, "服务内部错误")
	ParamError               = NewErrNo(ParamErrorCode, "参数出错")
	AuthorizationFailedError = NewErrNo(AuthorizationFailedErrCode, "authorization failed")

	// Admin
	AdminExistedError         = NewErrNo(AdminExistedErrorCode, "用户已存在")
	AdminNotExistedError      = NewErrNo(AdminNotExistedErrorCode, "未找到该用户")
	UserNameOrPasswordError   = NewErrNo(UserNameOrPasswordErrorCode, "用户名或密码错误")
	PasswordNotCorrectError   = NewErrNo(PasswordNotCorrectErrorCode, "密码不正确")
	UserNameIsNotEnabledError = NewErrNo(UserNameIsNotEnabledErrorCode, "账号已被禁用")
)
