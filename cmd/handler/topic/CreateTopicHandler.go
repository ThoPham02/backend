package topic

import (
	"back-end/cmd/logic/topic"
	"back-end/cmd/svc"
	"back-end/core/http_response"
	"back-end/core/logger"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateTopicHandler godoc
// @Summary create_topic
// @Schemes
// @Description create_topic
// @Tags topic
// @Param CreateTopic body types.CreateTopicRequest false "create topic"
// @Success 200 {object} types.CreateTopicResponse
// @Failure 401
// @Security ApiKeyAuth
// @Router /nckh/lecturer/create_topic [POST]
func CreateTopicHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add trace_id to context
		ctx := context.WithValue(c.Request.Context(), "trace_id", logger.GenerateTraceID("create-topic-api"))
		// Init log helper with context (have trace_id)
		logHelper := logger.NewContextLog(ctx)
		// New object logic (all logic code will implement in this object)
		topViewDeletedByIDLogic := topic.NewTopicGetAllLogic(ctx, svcCtx, logHelper)

		// Call functions in logic to process

		result, err := topViewDeletedByIDLogic.TopicGetAll()
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, result)
	}
}
