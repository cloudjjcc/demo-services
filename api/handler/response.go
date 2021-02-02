package handler

import "github.com/gin-gonic/gin"

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func ResponseWithErr(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(200, Response{
		Code: 4000,
		Msg:  err.Error(),
	})
}

func ResponseWithData(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, Response{
		Code: 2000,
		Data: data,
	})
}
