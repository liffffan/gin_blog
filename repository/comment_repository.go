package repository

import (
	"fmt"
	"gin_study_blog/model"
	"time"
)

func GetCommentList(articleId int64, pageNum, pageSize int) (commentList []*model.Comment, err error) {
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d\n", pageNum, pageSize)
	}

	sqlstr := "select id, content, username, create_time from comment where article_id = ? and status = 1 order by create_time desc limit ?, ?"
	err = DB.Select(&commentList, sqlstr, articleId, pageNum, pageSize)
	return
}

func InsertComment(Comment *model.Comment) (err error) {
	if Comment == nil {
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

	sqlstr := "insert into comment (article_id, username,content,create_time, status) values(?,?,?,?,?)"
	_, err = DB.Exec(sqlstr, Comment.ArticleId, Comment.Username, Comment.Content, time.Now(), 1)
	if err != nil {
		return
	}

	sqlstr = "update article set comment_count = comment_count + 1 where id = ?"
	_, err = DB.Exec(sqlstr, Comment.ArticleId)
	if err != nil {
		return
	}

	return
}
