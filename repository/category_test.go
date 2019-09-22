package repository

import "testing"

func init() {
	dns := "root:12345678@tcp(localhost:3306)/gin_blog?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestGetCategoryList(t *testing.T) {
	var categoryIds []int64
	categoryIds = append(categoryIds, 1, 2, 3)
	categoryList, err := GetCategoryList(categoryIds)
	if err != nil {
		t.Errorf("get category list failed, err:%v\n", err)
		return
	}

	if len(categoryList) != len(categoryIds) {
		t.Errorf("get category list failed, len of categoryList:%d, categoryIds len:%d\n", len(categoryList), len(categoryIds))
	}

	for _, v := range categoryList {
		t.Logf("id:%d, category:%#v\n", v.CategoryId, v)
	}

}
