/*
  author='du'
  date='2020/1/23 14:38'
*/
package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

const keyRequestId = "requestId"

func main() {
	r := gin.Default()

	//现在想纪录path
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	//加个日志的中间件
	r.Use(func(context *gin.Context) {
		startTime := time.Now()
		context.Next() //写完了，你就继续执行吧。
		logger.Info("请求：",
			zap.String("path", context.Request.URL.Path),
			zap.Int("status", context.Writer.Size()),
			zap.Duration("duration", time.Now().Sub(startTime)))

	}, func(context *gin.Context) {
		context.Set(keyRequestId, rand.Int())
		context.Next()
	})

	r.GET("/ping", func(context *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if rid, exist := context.Get(keyRequestId); exist {
			h[keyRequestId] = rid //这个是键值对存储的
		}
		context.JSON(200, h)
	})

	r.GET("/hb", func(context *gin.Context) {
		context.String(200, "dw")
	})

	r.Run()
}
