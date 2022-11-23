package errcode

var (
	ErrorGetTagListFail    = InitError(20010001, "获取标签列表失败")
	ErrorCreateTagFail     = InitError(20010002, "创建标签失败")
	ErrorUpdateTagFail     = InitError(20010003, "更新标签失败")
	ErrorDeleteTagFail     = InitError(20010004, "删除标签失败")
	ErrorCountTagFail      = InitError(20010005, "统计标签失败")
	TagNotFound            = InitError(40010006, "标签ID不存在")
	ErrorCreateArticleFail = InitError(20020001, "创建文章失败")
)
