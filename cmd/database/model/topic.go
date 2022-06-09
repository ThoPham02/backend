package model

type Topic struct {
	TopicID        int64  `db:"id"`
	TopicName      string `db:"name"`
	GroupStudentID int64  `db:"group_student_id"`
	LecturerID     int64  `db:"lecturer_id"`
	StartDay       string `db:"start_day"`
	EndDay         string `db:"end_day"`
	Allowance      string `db:"allowance"`
	TopicStatus    string `db:"status"`
}
