package response

import (
	"net/http"
	"os"

	"github.com/Jasonbourne723/socrates/global"
	"github.com/gin-gonic/gin"
)

type Response[T interface{}] struct {
	ErrorCode int    `json:"error_code"`
	Data      T      `json:"data"`
	Message   string `json:"message"`
}

func ServerError(c *gin.Context, err interface{}) {
	msg := "Internal Server Error"
	if global.App.Config.App.Env != "production" && os.Getenv(gin.EnvGinMode) != gin.ReleaseMode {
		if _, ok := err.(error); ok {
			msg = err.(error).Error()
		}
	}
	c.JSON(http.StatusInternalServerError, Response[any]{
		http.StatusInternalServerError,
		nil,
		msg,
	})
	c.Abort()
}

func Success[T interface{}](c *gin.Context, data T) {
	c.JSON(http.StatusOK, Response[T]{
		0,
		data,
		"ok",
	})
}

func Fail(c *gin.Context, errorCode int, msg string) {
	c.JSON(http.StatusOK, Response[any]{
		errorCode,
		nil,
		msg,
	})
}

func FailByError(c *gin.Context, error global.CustomError) {
	Fail(c, error.ErrorCode, error.ErrorMsg)
}

func ValidateFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.ValidateError.ErrorCode, msg)
}

func BusinessFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.BusinessError.ErrorCode, msg)
}

func TokenFail(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response[any]{
		global.Errors.TokenError.ErrorCode,
		nil,
		"无权限操作",
	})
}
