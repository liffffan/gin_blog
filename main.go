package main

import (
	"gin_study_blog/controller"
	"gin_study_blog/repository"
	"github.com/gin-gonic/gin"
)

// database/sql 里官方已经实现了连接池

func main() {
	router := gin.Default()

	dns := "root:root@tcp(localhost:3306)/gin_blog?parseTime=true"
	err := repository.Init(dns)
	if err != nil {
		panic(err)
	}

	// 静态文件目录
	router.Static("/static/", "./static")

	// 加载模版文件
	router.LoadHTMLGlob("views/*")

	// 首页
	router.GET("/", controller.IndexHandle)

	// 发布文章页面
	router.GET("/article/new/", controller.NewArticle)

	// 文章提交接口
	router.POST("/article/submit/", controller.ArticleSubmit)

	// 文章详情页
	router.GET("/article/detail/", controller.ArticleDetail)

	// 提交评论
	router.POST("/comment/submit/", controller.CommentSubmit)

	// 留言页面
	router.GET("/leave/new/", controller.LeaveNew)

	// 提交留言
	router.POST("/leave/submit/", controller.LeaveSubmit)

	// 分类云文章
	router.GET("/category/", controller.CategoryList)
	/*
		// 文章上传接口
		router.POST("/upload/file/", controller.UploadFile)


		// 关于我页面
		router.GET("/about/me/", controller.AboutMe)
	*/

	router.Run(":8080")

}
