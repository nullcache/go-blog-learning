package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
	validator "github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

// 原本是作为中间件，但是会意外反复注册，实际上注册一次就够了
var trans ut.Translator

func init() {
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		_ = zh_translations.RegisterDefaultTranslations(v, trans)
	}
}

type ValidError struct {
	Key     string
	Message string
}

type ValidErrorList []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrorList) Error() string {
	return strings.Join(v.ToErrorList(), ",")
}

func (v ValidErrorList) ToErrorList() []string {
	var errs = make([]string, 0, 5)
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrorList) {
	var errs ValidErrorList
	err := c.ShouldBind(v)
	if err != nil {
		verrs, ok := err.(val.ValidationErrors)
		if !ok {
			return false, errs // 这里errs是空切片啊
		}

		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}

		return false, errs
	}

	return true, nil
}
