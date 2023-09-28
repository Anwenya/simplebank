package api

import (
	"com.wlq/simplebank/util"
	"github.com/go-playground/validator/v10"
)

// 自定义参数校验标签 用于验证货币参数是否是支持范围内
var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportCurrency(currency)
	}
	return false
}
