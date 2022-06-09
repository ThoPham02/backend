package model

type Login struct {
	Mail         string `db:"mail"`
	Password     string `db:"password"`
	ActivateFlag int64  `db:"activate_flag"`
	LoginType    int64  `db:"login_type"`
}
