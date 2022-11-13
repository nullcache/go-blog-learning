package errcode

var (
	Success                   = InitError(0, "成功")
	ServerError               = InitError(9999, "服务内部错误")
	InvalidParams             = InitError(1001, "入参错误")
	NotFound                  = InitError(1002, "找不到")
	UnauthorizedAuthNotExist  = InitError(1003, "鉴权失败，找不到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = InitError(1004, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout  = InitError(1005, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate = InitError(1006, "鉴权失败，Token 生成失败")
	TooManyRequests           = InitError(1007, "请求过多")
)
