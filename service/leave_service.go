package service

import (
	"fmt"
	"gin_study_blog/model"
	"gin_study_blog/repository"
)

func GetLeaveList() (leaveList []*model.Leave, err error) {

	/*
		if articleId < 0 {
			err = fmt.Errorf("invalid parameter, articleId:%d", articleId)
			return
		}
		commentList, err = repository.GetCommentList(articleId)
		if err != nil {
			fmt.Printf("1 get comment list failed, err:%v\n", err)
			return
		}
		return
	*/

	// 调用 GetLeaveList 获取留言列表
	leaveList, err = repository.GetLeaveList(0, 15)
	return
}

func InsertLeave(author, content, email string) (err error) {
	Leave := &model.Leave{}
	Leave.UserName = author
	Leave.Content = content
	Leave.Email = email

	err = repository.InsertLeave(Leave)
	if err != nil {
		err = fmt.Errorf("insert leave failed, err:%v\n", err)
		return
	}
	return
}
