package errcode

var (
	Success                   = InitError(0, "成功")
	ServerError               = InitError(9999, "服务内部错误")
	InvalidParams             = InitError(1001, "入参错误")
	NotFound                  = InitError(1002, "错误的路由或请求方法")
	UnauthorizedAuthNotExist  = InitError(1003, "用户名或密码错误")
	EmptyTokenError           = InitError(1004, "未携带token")
	UnauthorizedTokenError    = InitError(1005, "token错误")
	UnauthorizedTokenTimeout  = InitError(1006, "token超时")
	UnauthorizedTokenGenerate = InitError(1007, "token生成失败")
	TooManyRequests           = InitError(1008, "请求过多")
	RequestTimeout            = InitError(1009, "请求超时")
)
