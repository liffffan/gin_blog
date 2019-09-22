package repository

import (
	"testing"
)

func init() {
	dns := "root:12345678@tcp(localhost:3306)/gin_blog?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

//func TestInsertArticle(t *testing.T) {
//	article := &model.ArticleDetail{}
//	article.ArticleInfo.CategoryId = 1
//	article.ArticleInfo.CommentCount = 0
//	article.ArticleInfo.CreateTime = time.Now()
//	article.ArticleInfo.Summary = "文章摘要牛皮哦"
//	article.ArticleInfo.Title = "Golang 牛皮"
//	article.ArticleInfo.Username = "golang"
//	article.ArticleInfo.ViewCount = 1
//	article.ArticleInfo.CommentCount = 1
//	article.Category.CategoryId = 1
//	article.Content = "大连桑达科技大厦路口就到拉空间打开链接的考拉三等奖考拉三等奖askedtalked "
//
//	articleId, err := InsertArticle(article)
//	if err != nil {
//		t.Errorf("insert article failed, err:%v\n", err)
//		return
//	}
//
//	t.Logf("insert article success, articleId:%d\n", articleId)
//
//
//
//}

func TestGetArticleList(t *testing.T) {
	articleList, err := GetArticleList(1, 15)
	if err != nil {
		t.Errorf("get article failed, err:%v\n", err)
		return
	}
	t.Logf("get article success, len:%d\n", len(articleList))

}
