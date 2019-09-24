package model

import "time"

type Leave struct {
	Id         int64     `db:"id"`
	UserName   string    `db:"username"`
	Email      string    `db:"email"`
	Content    string    `db:"content"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}
