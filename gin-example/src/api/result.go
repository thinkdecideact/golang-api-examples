package api

import (
	"github.com/gin-gonic/gin"
)

const SUCCESS_CODE = 0
const SUCCESS_MSG = "Success"

const FAILURE_CODE = 1
const FAILURE_MSG = "Failure"

const TOKEN_INVALID_CODE = -1
const TOKEN_INVALID_MSG = "token expired, login required"

func Result(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func Success(ctx *gin.Context, msg string, optionalDataList ...interface{}) {
	var data interface{}
	if len(optionalDataList) > 0 {
		data = optionalDataList[0]
	}
	Result(ctx, SUCCESS_CODE, msg, data)
}

func Failure(ctx *gin.Context, msg string, optionalDataList ...interface{}) {
	var data interface{}
	if len(optionalDataList) > 0 {
		data = optionalDataList[0]
	}
	Result(ctx, FAILURE_CODE, msg, data)
}
