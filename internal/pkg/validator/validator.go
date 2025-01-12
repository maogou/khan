package validator

import (
	"errors"

	"github.com/gin-gonic/gin/binding"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

// init  注册中文翻译器
func init() {
	//注册中文翻译器
	zhTrans := zh.New()
	uni = ut.New(zhTrans, zhTrans)

	// 获取gin的验证器,不要验证错误了,因为gin默认用的就是这个验证器
	validate = binding.Validator.Engine().(*validator.Validate)

	trans, _ = uni.GetTranslator("zh")

	_ = zh_translations.RegisterDefaultTranslations(validate, trans)
}

// Translate  翻译错误信息并返回一个错误
func Translate(errs error) map[string]string {
	res := make(map[string]string, 0)

	var verse validator.ValidationErrors
	ok := errors.As(errs, &verse)

	if ok {
		for _, e := range verse {
			res[e.StructField()] = e.Translate(trans)
		}
	} else {
		res["validate"] = "传递参类型不匹配,参数绑定错误"
	}

	return res
}
