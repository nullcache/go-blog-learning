package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(9999, "服务内部错误")
	InvalidParams             = NewError(1001, "入参错误")
	NotFound                  = NewError(1002, "找不到")
	UnauthorizedAuthNotExist  = NewError(1003, "鉴权失败，找不到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewError(1004, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout  = NewError(1005, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate = NewError(1006, "鉴权失败，Token 生成失败")
	TooManyRequests           = NewError(1007, "请求过多")
)
