package repository

import (
	"blog/model"
)

func GetCommentList(articleId int64, pageNum, pageSize int) (commentList []*model.Comment, err error) {
	sqlstr := "select id, content, username, create_time from comment wher article_id = ? and status = 1 order by create_time desc limit ?, ?"
	err = DB.Select(&commentList, sqlstr, articleId, pageNum, pageSize)
	return
}
