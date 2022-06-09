package model

type Student struct {
	Id           int64  `db:"id"`
	Name         string `db:"name"`
	DepartmentId int64  `db:"department_id"`
	UserType     int64  `db:"user_type"`
}
