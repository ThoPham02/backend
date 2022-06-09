package handler

import (
	"back-end/cmd/handler/topic"
	"back-end/cmd/svc"
	"back-end/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterHandlers(router *gin.Engine, serverCtx *svc.ServiceContext) {
	docs.SwaggerInfo.BasePath = ""
	router.LoadHTMLGlob("cmd/templates/pages/*.tmpl")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/nckh/topic", topic.GetAllTopicsHandler(serverCtx))
	router.GET("/nckh/search_topic_by_key", topic.GetTopicByKeyHandler(serverCtx))
	router.POST("/nckh/lecturer/create_topic", topic.CreateTopicHandler(serverCtx))
	router.PUT("/nckh/lecturer/topic/{id}", topic.TopicUpdateHandler(serverCtx))
}
