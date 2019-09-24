package repository

import (
	"database/sql"
	"fmt"
	"gin_study_blog/model"
	"time"
)

func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {

	if article == nil {
		err = fmt.Errorf("invalid article parameter")
		return
	}

	sqlstr := "insert into article(content, summary, title, username, category_id, view_count, comment_count, create_time) value(?,?,?,?,?,?,?,?)"
	result, err := DB.Exec(sqlstr, article.Content, article.Summary,
		article.Title, article.Username, article.ArticleInfo.CategoryId, article.ViewCount, article.CommentCount, time.Now())
	if err != nil {
		return
	}

	articleId, err = result.LastInsertId()
	return
}

func GetArticleList(pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d", pageNum, pageSize)
		return
	}

	sqlstr := `select id, category_id, summary, title, view_count, create_time, comment_count, username from article where status = 1 order by create_time desc limit ?,?`

	err = DB.Select(&articleList, sqlstr, pageNum, pageSize)
	return
}

func GetRelativeArticle(articleId int64) (articleList []*model.RelativeArticle, err error) {
	var categoryId int64
	sqlstr := "select category_id from article where id=?"
	err = DB.Get(&categoryId, sqlstr, articleId)
	if err != nil {
		return
	}

	sqlstr = "select id, title from article where category_id=? and id !=?  limit 10"
	err = DB.Select(&articleList, sqlstr, categoryId, articleId)
	return
}

func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	sqlstr := "select id, category_id, summary, title, view_count, create_time, comment_count, username,category_id, content from article where id = ? and status = 1"
	articleDetail = &model.ArticleDetail{}
	err = DB.Get(articleDetail, sqlstr, articleId)
	if err != nil {
		err = fmt.Errorf("get article failed, err:%v\n", err)
		return
	}
	return
}

func GetPrevArticleById(articleId int64) (prevArticle *model.RelativeArticle, err error) {
	/*
		sqlstr := "select id, title from article where id = ?"
		err = DB.Get(&prevArticle, sqlstr, articleId - 1)
		if err == sql.ErrNoRows {
			prevArticle.ArticleId = 0
			prevArticle.Title = ""
			return
		}
		return

	*/

	prevArticle = &model.RelativeArticle{
		ArticleId: -1,
	}
	sqlstr := "select id, title from article where id < ? and status = 1 order by id desc limit 1"
	err = DB.Get(prevArticle, sqlstr, articleId)
	if err != nil {
		return
	}

	return
}

func GetNextArticleById(articleId int64) (nextArticle *model.RelativeArticle, err error) {
	nextArticle = &model.RelativeArticle{
		ArticleId: -1,
	}
	sqlstr := "select id, title from article where id > ? and status = 1 order by id asc limit 1"
	err = DB.Get(nextArticle, sqlstr, articleId)
	if err != nil {
		return
	}

	return
}

func IsArticleExist(articleId int64) (exists bool, err error) {

	var id int64
	sqlstr := "select id from article where id = ?"
	err = DB.Get(&id, sqlstr, articleId)
	// 查询一条记录时, 不能使用类似if err := db.QueryRow().Scan(&...); err != nil {}的处理方式
	// 因为查询单条数据时, 可能返回var ErrNoRows = errors.New("sql: no rows in result set")该种错误信息
	// 而这属于正常错误
	if err == sql.ErrNoRows {
		exists = false
		return
	}

	if err != nil {
		return
	}

	exists = true
	return
}

func GetCategoryArticle(categoryId int64, pageNum, pageSize int) (categoryArticleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d", pageNum, pageSize)
		return
	}

	sqlstr := `select id, category_id, summary, title, view_count, create_time, comment_count, username from article where status = 1 and category_id = ? order by create_time desc limit ?,?`
	err = DB.Select(&categoryArticleList, sqlstr, categoryId, pageNum, pageSize)
	return
}
