package types

type TopicGetBySearchingRequest struct {
	Key string `form:"key" json:"key"`
}

type TopicGetBySearchingResponse struct {
	TopicResponse []Topic `json:"topic_response"`
}
