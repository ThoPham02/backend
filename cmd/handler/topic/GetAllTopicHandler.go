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

// GetAllTopicsHandler godoc
// @Summary get_all
// @Schemes
// @Description get_all
// @Tags topic
// @Success 200 {object} types.TopicGetAllResponse
// @Failure 401 {object} types.HttpErrorResponse
// @Security ApiKeyAuth
// @Router /nckh/topic [GET]
func GetAllTopicsHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add trace_id to context
		ctx := context.WithValue(c.Request.Context(), "trace_id", logger.GenerateTraceID("get-all-topic-api"))
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
