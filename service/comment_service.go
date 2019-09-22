package service

import (
	"blog/model"
	"blog/repository"
	"fmt"
)

func GetCommentList(articleId int64) (commentList []*model.Comment, err error) {

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

	//1. 首先，要验证article_id是否合法
	exist, err := repository.IsArticleExist(articleId)
	if err != nil {
		fmt.Printf("query database failed, err:%v\n", err)
		return
	}

	if exist == false {
		err = fmt.Errorf("article id:%d not found", articleId)
		return
	}

	//2. 调用dal GetCommentList获取评论列表
	commentList, err = repository.GetCommentList(articleId, 0, 100)
	return
}
