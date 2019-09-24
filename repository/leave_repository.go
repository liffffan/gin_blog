package repository

import (
	"fmt"
	"gin_study_blog/model"
	"time"
)

func GetLeaveList(pageNum, pageSize int) (commentList []*model.Leave, err error) {
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d\n", pageNum, pageSize)
	}

	sqlstr := "select id, content, username, email, create_time from `leave` order by create_time desc limit ?, ?"
	err = DB.Select(&commentList, sqlstr, pageNum, pageSize)
	return
}

func InsertLeave(Leave *model.Leave) (err error) {
	if Leave == nil {
		err = fmt.Errorf("invalid parameter")
		return
	}

	tx, err := DB.Beginx()
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	sqlstr := "insert into `leave` (username,email,content,create_time) values(?,?,?,?)"
	_, err = DB.Exec(sqlstr, Leave.UserName, Leave.Email, Leave.Content, time.Now())
	if err != nil {
		return
	}

	return
}
