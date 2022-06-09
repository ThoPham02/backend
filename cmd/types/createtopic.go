package types

type CreateTopicRequest struct {
	TopicName      string `json:"name"`
	GroupStudentID int64  `json:"group_student_id"`
	LecturerID     int64  `json:"lecturer_id"`
	StartDay       string `json:"start_day"`
	EndDay         string `json:"end_day"`
	Allowance      string `json:"allowance"`
	TopicStatus    string `json:"status"`
}

type CreateTopicResponse struct {
}
