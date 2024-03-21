package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.GET("")
		userGroup.PUT("")
		userGroup.POST("")
		userGroup.DELETE("")
	}
	return r
}

func statCost() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		time.Sleep(1000)
		context.Next()
		cost := time.Since(start)
		log.Printf("cost time: [%#v]", cost)
	}
}
