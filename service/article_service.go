package service

import (
	"fmt"
	"gin_study_blog/model"
	"gin_study_blog/repository"
	"math"
)

func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64) {
LABEL:
	// 这里进行了去重，比如10篇文章的 id 都是 1，那么 ids 里就是 10 个 1，所以要去重
	for _, article := range articleInfoList {
		categoryId := article.CategoryId
		for _, id := range ids {
			// 如果里面已经存在这个 id 了，就跳过这次循环
			if id == categoryId {
				continue LABEL
			}
		}
		ids = append(ids, categoryId)
	}
	return
}

func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	// 1.从数据库中，获取文章列表
	articleInfoList, err := repository.GetArticleList(pageNum, pageSize)
	if err != nil {
		fmt.Printf("1 get article list failed, err:%v\n", err)
		return
	}

	if len(articleInfoList) == 0 {
		return
	}

	// 拿到文章所有分类 id 的列表
	categoryIds := getCategoryIds(articleInfoList)
	// 2.从数据库中，获取文章对应的分类信息
	categoryList, err := repository.GetCategoryList(categoryIds)
	if err != nil {
		fmt.Printf("2 get category list failed, err:%v\n", err)
		return
	}

	// 聚合数据
	for _, article := range articleInfoList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		categoryId := article.CategoryId
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}

		articleRecordList = append(articleRecordList, articleRecord)

	}
	return
}

func InsertArticle(content, author, title string, categoryId int64) (err error) {
	// 构造一个 ArticleDetail 的结构体
	//var article *model.ArticleDetail
	article := &model.ArticleDetail{}
	article.Content = content
	article.ArticleInfo.Username = author
	article.ArticleInfo.CategoryId = categoryId
	article.Title = title
	article.CommentCount = 0
	article.ViewCount = 0
	// 转成 UTF-8 切片后截取128个字节
	// 如果文章小于128字节可能会越界，所以定义一个最小值
	contentUtf8 := []rune(content)
	minLength := int(math.Min(float64(len(contentUtf8)), 128.0))
	article.Summary = string([]rune(content)[:minLength])

	id, err := repository.InsertArticle(article)
	if err != nil {
		fmt.Printf("insert article failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert artilce success, id:%d, err:%v\n", id, err)

	return nil
}

func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {

	articleDetail, err = repository.GetArticleDetail(articleId)
	if err != nil {
		fmt.Printf("get article detail failed, err:%v\n", err)
		return
	}
	return
}

func GetRelativeArticleList(articleId int64) (articleList []*model.RelativeArticle, err error) {
	if articleId < 0 {
		err = fmt.Errorf("invalid parameter, articleId:%d", articleId)
		return
	}

	articleList, err = repository.GetRelativeArticle(articleId)
	if err != nil {
		fmt.Printf("get relative article list failed, err:%v\n", err)
		return
	}

	return

}

func GetPrevAndNextArticleInfo(articleId int64) (prevArticle, nextArticle *model.RelativeArticle, err error) {
	if articleId < 0 {
		err = fmt.Errorf("invalid parameter, articleId:%d", articleId)
		return
	}

	prevArticle, err = repository.GetPrevArticleById(articleId)
	if err != nil {
		fmt.Printf("get prev article failed, err:%v\n", err)
		return
	}

	nextArticle, err = repository.GetNextArticleById(articleId)
	if err != nil {
		fmt.Printf("get pre article failed, err:%v\n", err)
		return
	}

	return

}
