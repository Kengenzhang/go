package v1

import (
	"aa/model"
	"aa/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// //查询用户是否存在
// func UserExist(c *gin.Context)  {

// }
// 添加用户
var r_code int

// 接收从user模块传过来的的code
func AddUser(c *gin.Context) {
	var data model.User

	c.ShouldBindJSON(&data)
	r_code = model.CheckUser(data.Username)
	//  接收查询的返回值
	if r_code == errmsg.SUCCSE {
		// 成功的话将数据写入
		model.CreateUser(&data)
	}
	if r_code == errmsg.ERROR_USERNAME {
		//1001,用户存在
		r_code = errmsg.ERROR_USERNAME
	}
	c.JSON(http.StatusOK, gin.H{

		"status":  r_code,
		"data":    data,
		"message": errmsg.Geterrmsg(r_code),
	})

}

//查询用户

// 查询用户列表
func GetUsers(c *gin.Context) {
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
	data := model.Getusers(pagesize, pagenum)
	c.JSON(http.StatusOK, gin.H{
		"status":  r_code,
		"data":    data,
		"message": errmsg.Geterrmsg(r_code),
	})
}

// 编辑用户
func Edituser(c *gin.Context) {
 var data model.User
 id, _ := strconv.Atoi(c.Param("id"))
//  获取id
 c.ShouldBindJSON(&data)
 r_code=model.CheckUser(data.Username)
//  查询是否重名
 if r_code!=errmsg.ERROR {
	model.Ususer(id,&data)
	// 不重名返回给ususer方法进行编辑处理
 }
 if r_code == errmsg.ERROR_USERNAME {
	//1001,用户存在
	r_code = errmsg.ERROR_USERNAME
}
c.JSON(http.StatusOK, gin.H{

	"status":  r_code,
	"message": errmsg.Geterrmsg(r_code),
})
}
// 删除用户(软删除)
func DelUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// 将前端传过来的数据进行性格式转换
	r_code = model.Deluser(id)
	// 节后数据库操作之后的值传递
	c.JSON(http.StatusOK, gin.H{
		"status":  r_code,
		"message": errmsg.Geterrmsg(r_code),
	})
}
