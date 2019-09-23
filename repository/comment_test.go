package repository

import (
	"fmt"
	"gin_study_blog/model"
	"testing"
)

func TestGetCommentList(t *testing.T) {
	CommentList, err := GetCommentList(5, 0, 15)
	if err != nil {
		t.Errorf("insert comment failed, err:%v\n", err)
		return
	}

	for _, v := range CommentList {
		fmt.Printf("Comment:%#v\n", v)
	}
	t.Logf("get article success, len:%d\n", len(CommentList))

}

func TestInsertComment(t *testing.T) {
	Comment := &model.Comment{}
	Comment.Username = "李翔"
	Comment.Content = "达会计师的骄傲开始登记卡萨大家开始登记哈三等奖库函数"
	Comment.ArticleId = 8

	err := InsertComment(Comment)
	if err != nil {
		t.Errorf("insert comment failed, err:%v\n", err)
		return
	}

}
