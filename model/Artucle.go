package model

// 文章模块
import (
	"aa/utils/errmsg"

	"github.com/jinzhu/gorm"
)

type Artucle struct {
	gorm.Model
	Category Category `gorm:"foreignkey:Cid"`
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Cid      int      `gorm:"type:int;not null" json:"cid"`
	Desc     string   `gorm:"type:varchar(100)" json:"desc"`
	Context  string   `gorm:"type:longtext" json:"context"`
	Img      string   `gorm:"type:varchar(100)" json:"img"`
}

// 新增文章
func CreateArtucle(data *Artucle) int {
	err2 := db.Create(&data).Error
	if err2 != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE
}

// 查询文章列表
func GetArtucle(pagesize int, pagenum int) ([]Artucle, int) {
	var artlist []Artucle
	err2 := db.Preload("Category").Limit(pagesize).Offset((pagenum - 1) * pagesize).Find(&artlist).Error
	if err2 != nil && err2 != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return artlist, errmsg.SUCCSE
}

// 删除文章
func DelArtucle(id int) int {
	var art Artucle
	err2 := db.Where("id=?", id).Delete(&art).Error
	if err2 != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 编辑文章
func UsArtucle(id int, data *Artucle) int {
	var art Artucle
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["context"] = data.Context
	maps["img"] = data.Img
	err2 := db.Model(&art).Where("id=?", id).Update(maps).Error
	if err2 != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE

}

// 查询单个文章
func GETartinfo(id int) (Artucle, int) {
	var art Artucle
	err2 := db.Preload("Category").Where("id=?", id).First(&art).Error
	if err2 != nil {
		return art, errmsg.ERROR_ARTNOT
	}
	return art, errmsg.SUCCSE
}

// 查询分类下所有文章
func GETcateallart(id int, pagesize int, pagenum int) ([]Artucle, int) {
	var cateart []Artucle
	err := db.Preload("Category").Limit(pagesize).Offset((pagenum-1)*pagesize).Where("id=?", id).Find(&cateart).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST
	}
	return cateart, errmsg.SUCCSE
}
