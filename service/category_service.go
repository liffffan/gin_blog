package service

import (
	"fmt"
	"gin_study_blog/model"
	"gin_study_blog/repository"
)

func GetAllCategoryList() (categoryList []*model.Category, err error) {
	//1. 从数据库中，获取文章分类列表
	categoryList, err = repository.GetAllCategoryList()
	if err != nil {
		fmt.Printf("1 get category list failed, err:%v\n", err)
		return
	}

	return
}
