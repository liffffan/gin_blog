package repository

import (
	"gin_study_blog/model"
	"time"
)

func GetCommentList(articleId int64, pageNum, pageSize int) (commentList []*model.Comment, err error) {
	sqlstr := "select id, content, username, create_time from comment where article_id = ? and status = 1 order by create_time desc limit ?, ?"
	err = DB.Select(&commentList, sqlstr, articleId, pageNum, pageSize)
	return
}

func InsertComment(Comment *model.Comment) (err error) {
	sqlstr := "insert into comment (article_id, username,content,create_time, status) values(?,?,?,?,?)"
	_, err = DB.Exec(sqlstr, Comment.ArticleId, Comment.Username, Comment.Content, time.Now(), 1)
	if err != nil {
		return
	}
	return
}
