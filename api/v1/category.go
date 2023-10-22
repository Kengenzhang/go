package v1

import (
	"aa/model"
	"aa/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 接收从catgory模块传过来的的code
func Addcategory(c *gin.Context) {
	var data model.Category

	c.ShouldBindJSON(&data)
	r_code = model.Checkcatgory(data.Name)
	//  接收查询的返回值
	if r_code == errmsg.SUCCSE {
		// 成功的话将数据写入
		model.Createcategory(&data)
	}
	if r_code == errmsg.ERROR_CATEGORY {
		//1001,
		r_code = errmsg.ERROR_CATEGORY
	}
	c.JSON(http.StatusOK, gin.H{

		"status":  r_code,
		"data":    data,
		"message": errmsg.Geterrmsg(r_code),
	})

}

// 查询分类列表
func Getcategory(c *gin.Context) {
	//pagesize和pagenum接收来自前端的数据并通过atio方法转换成int型
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	pagenum, _ := strconv.Atoi(c.Query("pagenum"))
	if pagesize == 0 {
		pagesize = -1
		// 不查询的话等于-1
	}
	if pagenum == 0 {
		pagenum = -1
	}
	r_code = errmsg.SUCCSE
	data := model.Getcategory(pagesize, pagenum)
	c.JSON(http.StatusOK, gin.H{
		"status":  r_code,
		"data":    data,
		"message": errmsg.Geterrmsg(r_code),
	})
}

// 编辑分类
func Editcategory(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	//  获取id
	c.ShouldBindJSON(&data)
	r_code = model.Checkcatgory(data.Name)
	//  查询是否重名
	if r_code != errmsg.ERROR {
		model.Uscategory(id, &data)
		// 不重名返回给ususer方法进行编辑处理
	}
	if r_code == errmsg.ERROR_CATEGORY {
		//1001,用户存在
		r_code = errmsg.ERROR_CATEGORY
	}
	c.JSON(http.StatusOK, gin.H{

		"status":  r_code,
		"message": errmsg.Geterrmsg(r_code),
	})

}

// 删除分类(软删除)
func Delcategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// 将前端传过来的数据进行性格式转换
	r_code = model.Delcategory(id)
	// 节后数据库操作之后的值传递
	c.JSON(http.StatusOK, gin.H{
		"status":  r_code,
		"message": errmsg.Geterrmsg(r_code),
	})
}
