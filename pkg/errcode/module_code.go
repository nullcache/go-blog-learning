package errcode

var (
	ErrorGetTagListFail     = InitError(20010001, "获取标签列表失败")
	ErrorCreateTagFail      = InitError(20010002, "创建标签失败")
	ErrorUpdateTagFail      = InitError(20010003, "更新标签失败")
	ErrorDeleteTagFail      = InitError(20010004, "删除标签失败")
	ErrorCountTagFail       = InitError(20010005, "统计标签失败")
	TagNotFound             = InitError(20010006, "标签ID不存在")
	ErrorCreateArticleFail  = InitError(20020001, "创建文章失败")
	ArticleNotFound         = InitError(20020002, "文章ID不存在")
	ErrorGetArticleListFail = InitError(20020003, "获取标签列表失败")
	ErrorDelArticleFail     = InitError(20020004, "删除标签失败")
	ErrorUpdateArticleFail  = InitError(20020005, "更新标签失败")
)
