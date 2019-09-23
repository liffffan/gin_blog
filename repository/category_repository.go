package repository

import (
	"gin_study_blog/model"
	"github.com/jmoiron/sqlx"
)

func GetCategoryList(categoryIds []int64) (categoryList []*model.Category, err error) {
	sqlstr, args, err := sqlx.In("select id, category_name, category_no from category where id in (?)", categoryIds)
	if err != nil {
		return
	}

	err = DB.Select(&categoryList, sqlstr, args...)
	return
}

func GetAllCategoryList() (categoryList []*model.Category, err error) {
	sqlstr := "select id, category_name, category_no from category order by category_no asc"
	err = DB.Select(&categoryList, sqlstr)
	return
}

func GetCategoryById(categoryId int64) (category *model.Category, err error) {
	sqlstr := "select id, category_name, category_no from category where id = ?"
	category = &model.Category{}
	err = DB.Get(category, sqlstr, categoryId)
	return
}
