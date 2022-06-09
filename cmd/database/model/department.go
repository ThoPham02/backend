package model

type Department struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}
