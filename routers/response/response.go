package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"unbound/pkg/setting"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	en "github.com/go-playground/locales/en"
	zh "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
	zhTranslation "github.com/go-playground/validator/v10/translations/zh"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var trans ut.Translator

func init() {
	// 获取当前设置的语言
	var lang = setting.Server.Lang

	// 如果语言设置为中文
	if lang == "zh" {
		// 创建中文翻译器
		trans, _ = ut.New(zh.New()).GetTranslator("zh")
		// 注册中文翻译到验证器
		if err := zhTranslation.RegisterDefaultTranslations(binding.Validator.Engine().(*validator.Validate), trans); err != nil {
			// 如果注册失败，则打印错误信息
			fmt.Println("validator zh translation error", err)
		}
	}

	// 如果语言设置为英文
	if lang == "en" {
		// 创建英文翻译器
		trans, _ = ut.New(en.New()).GetTranslator("en")
		// 注册英文翻译到验证器
		if err := enTranslation.RegisterDefaultTranslations(binding.Validator.Engine().(*validator.Validate), trans); err != nil {
			// 如果注册失败，则打印错误信息
			fmt.Println("validator en translation error", err)
		}
	}
}

const (
	SUCCESS        = 0
	ERROR          = 500
	INVALID_PARAMS = 400
)

var MsgFlags = map[int]string{
	SUCCESS:        "success",
	ERROR:          "error",
	INVALID_PARAMS: "请求参数错误",
}

// Result 函数用于封装 gin 框架的 JSON 响应
//
// 参数：
// gin：gin.Context 类型，表示 HTTP 请求的上下文对象
// code：int 类型，表示 HTTP 响应的状态码
// msg：string 类型，表示 HTTP 响应的消息
// data：interface{} 类型，表示 HTTP 响应的数据
func Result(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

// Ok 函数用于向客户端返回成功响应
// 参数c为gin.Context类型，表示HTTP请求上下文
// 参数data为interface{}类型，表示要返回给客户端的数据
func Ok(c *gin.Context, data interface{}) {
	Result(c, SUCCESS, MsgFlags[SUCCESS], data)
}

// Error 函数用于处理gin框架中的错误请求，返回错误信息给客户端
// 参数c是gin框架中的上下文对象
// 参数message是错误信息字符串
func Error(c *gin.Context, message string) {
	Result(c, ERROR, message, nil)
	fmt.Println(message)
}

// ErrorValidation 函数用于处理错误验证，返回错误信息给客户端并在控制台打印错误信息
// 参数：
//
//	c: gin框架的上下文对象
//	err: 错误信息
//
// 返回值：无
func ErrorValidation(c *gin.Context, err error) {
	var message string

	// 根据错误类型进行处理
	switch e := err.(type) {

	case validator.ValidationErrors:
		// 如果是验证错误，遍历错误列表并拼接错误信息
		for _, valErr := range e {
			message += valErr.Translate(trans) + "; "
		}

	case *json.UnmarshalTypeError:
		// 如果是 JSON 解析类型错误，构造错误信息
		message = fmt.Sprintf("参数 '%s' 类型不匹配，期望的类型是 '%v'", e.Field, e.Type)

	default:
		// 其他错误类型，构造未知错误信息
		message = "发生未知错误: " + err.Error()
	}

	message = "请求参数参数存在问题: " + message

	// 返回错误信息给客户端
	Result(c, INVALID_PARAMS, message, nil)

	// 打印错误信息到控制台
	fmt.Println(message)
}
