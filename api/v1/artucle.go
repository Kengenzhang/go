package v1

import (
	"aa/model"
	"aa/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Addartucle(c *gin.Context) {
	var data model.Artucle

	c.ShouldBindJSON(&data)
	r_code = model.CreateArtucle(&data)
	c.JSON(http.StatusOK, gin.H{

		"status":  r_code,
		"data":    data,
		"message": errmsg.Geterrmsg(r_code),
	})

}

// 查询分类下所有文章
func GETcateart(c *gin.Context) {
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	pagenum, _ := strconv.Atoi(c.Query("pagesize"))
	id, _ := strconv.Atoi(c.Param("id"))
	if pagesize == 0 {
		pagesize = -1
		// 不查询的话等于-1
	}
	if pagenum == 0 {
		pagenum = -1
	}
	data, r_code := model.GETcateallart(id, pagesize, pagenum)
	c.JSON(http.StatusOK, gin.H{
		"status":  r_code,
		"data":    data,
		"message": errmsg.Geterrmsg(r_code),
	})
}

// 查询单个文章
func GETartinfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, r_code := model.GETartinfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  r_code,
		"data":    data,
		"message": errmsg.Geterrmsg(r_code),
	})
}

// 查询文章列表
func Getartucle(c *gin.Context) {
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	pagenum, _ := strconv.Atoi(c.Query("pagenum"))
	if pagesize == 0 {
		pagesize = -1
	}
	if pagenum == 0 {
		pagenum = -1
	}
	data, r_code := model.GetArtucle(pagesize, pagenum)
	c.JSON(http.StatusOK, gin.H{
		"status":  r_code,
		"data":    data,
		"message": errmsg.Geterrmsg(r_code),
	})
}

// 编辑分类
func Editartucle(c *gin.Context) {
	var data model.Artucle
	id, _ := strconv.Atoi(c.Param("id"))
	//  获取id
	c.ShouldBindJSON(&data)
	r_code = model.UsArtucle(id, &data)
	c.JSON(http.StatusOK, gin.H{

		"status":  r_code,
		"message": errmsg.Geterrmsg(r_code),
	})

}

// 删除分类(软删除)
func Delartucle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// 将前端传过来的数据进行性格式转换
	r_code = model.DelArtucle(id)
	// 节后数据库操作之后的值传递
	c.JSON(http.StatusOK, gin.H{
		"status":  r_code,
		"message": errmsg.Geterrmsg(r_code),
	})
}
