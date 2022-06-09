package topic

import (
	"back-end/cmd/logic/topic"
	"back-end/cmd/svc"
	"back-end/cmd/types"
	"back-end/core/http_request"
	"back-end/core/http_response"
	"back-end/core/logger"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetTopicByKeyHandler godoc
// @Summary get_by_searching
// @Schemes
// @Description get_by_searching
// @Tags topic
// @Param Key query string true "Key to search"
// @Success 200 {object} types.TopicGetBySearchingResponse
// @Failure 401 {object} types.HttpErrorResponse
// @Security ApiKeyAuth
// @Router /nckh/search_topic_by_key [GET]
func GetTopicByKeyHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add trace_id to context
		ctx := context.WithValue(c.Request.Context(), "trace_id", logger.GenerateTraceID("get-by-searching-api"))
		// Init log helper with context (have trace_id)
		logHelper := logger.NewContextLog(ctx)
		// New object logic (all logic code will implement in this object)
		topicGetBySearchingLogic := topic.NewTopicGetBySearchingLogic(ctx, svcCtx, logHelper)
		request := &types.TopicGetBySearchingRequest{}
		err := http_request.BindQueryString(c, request)
		if err != nil {
			logHelper.Errorw("msg", "Failed while parse topic registration request", "error", err.Error())
			http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
			return
		}

		// Call functions in logic to process
		response, err := topicGetBySearchingLogic.TopicGetBySearching(request)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, response)
	}
}
