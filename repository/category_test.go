package repository

import "testing"

func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(1)
	if err != nil {
		t.Errorf("get category failed, err:%v\n", err)
		return
	}
	t.Logf("category:%#v\n", category)
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
