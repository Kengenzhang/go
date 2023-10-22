package model

import (
	"aa/utils/errmsg"

	"github.com/jinzhu/gorm"
)

// 分类模块

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}
// 查询分类
func Checkcatgory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name=?", name).First(&cate)
	if cate.ID > 0 {

		return errmsg.ERROR_CATEGORY 
	}
	return errmsg.SUCCSE
}

// 新增分类
func Createcategory(data *Category) int {
	err2 := db.Create(&data).Error
	if err2 != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE
}

// 查询分类列表
func Getcategory(pagesize int, pagenum int) []Category {
	var cate []Category
	err2 := db.Limit(pagesize).Offset((pagenum - 1) * pagesize).Find(&cate).Error
	if err2 != nil && err2 != gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}

// 删除分类
func Delcategory(id int) int {
	var cate Category
	err2 := db.Where("id=?", id).Delete(&cate).Error
	if err2 != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
// 编辑分类
func Uscategory(id int, data *Category) int {
	var cate Category
	
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err2 := db.Model(&cate).Where("id=?", id).Update(maps).Error
	if err2 != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}


/*
gorm.io/driver/mysql v1.5.1 h1:WUEH5VF9obL/lTtzjmML/5e6VfFR/788coz2uaVCAZw=
gorm.io/driver/mysql v1.5.1/go.mod h1:Jo3Xu7mMhCyj8dlrb3WoCaRd1FhsVh+yMXb1jUInf5o=
gorm.io/driver/sqlite v1.5.3 h1:7/0dUgX28KAcopdfbRWWl68Rflh6osa4rDh+m51KL2g=
gorm.io/driver/sqlite v1.5.3/go.mod h1:qxAuCol+2r6PannQDpOP1FP6ag3mKi4esLnB/jHed+4=
gorm.io/gorm v1.25.1/go.mod h1:L4uxeKpfBml98NYqVqwAdmV1a2nBtAec/cf3fpucW/k=
gorm.io/gorm v1.25.4 h1:iyNd8fNAe8W9dvtlgeRI5zSVZPsq3OpcTu37cYcpCmw=
gorm.io/gorm v1.25.4/go.mod h1:L4uxeKpfBml98NYqVqwAdmV1a2nBtAec/cf3fpucW/k=

*/
