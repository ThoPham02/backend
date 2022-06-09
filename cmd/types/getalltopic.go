package types

type Topic struct {
	TopicID        int64  `json:"id"`
	TopicName      string `json:"name"`
	GroupStudentID int64  `json:"group_student_id"`
	LecturerID     int64  `json:"lecturer_id"`
	StartDay       string `json:"start_day"`
	EndDay         string `json:"end_day"`
	Allowance      string `json:"allowance"`
	TopicStatus    string `json:"status"`
}

type TopicGetAllResponse struct {
	TopicResponse []Topic `json:"topic_response"`
}
