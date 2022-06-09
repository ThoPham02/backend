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

// TopicUpdateHandler godoc
// @Summary update_topic
// @Schemes
// @Description update_topic
// @Tags topic
// @Param id path string true "ID"
// @Param  Key body types.UpdateTopicRequest true "Key to search"
// @Success 200 {object} types.UpdateTopicInput
// @Failure 401 {object} types.HttpErrorResponse
// @Security ApiKeyAuth
// @Router /nckh/lecturer/topic/{id} [PUT]
func TopicUpdateHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "trace_id", logger.GenerateTraceID("update-topic-api"))
		logHelper := logger.NewContextLog(ctx)
		topicUpdateLogic := topic.NewTopicUpdateLogic(ctx, svcCtx, logHelper)
		path := &types.UpdateTopicPath{}
		err := http_request.BindUri(c, path)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
			return
		}

		request := &types.UpdateTopicRequest{}
		err = http_request.BindBodyJson(c, request)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
			return
		}

		input := &types.UpdateTopicInput{
			Path:    path,
			Request: request,
		}

		err = topicUpdateLogic.TopicUpdate(input)
		if err != nil {
			http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
			return
		}
		http_response.ResponseJSON(c, http.StatusOK, "Update Topic Success")
	}
}
