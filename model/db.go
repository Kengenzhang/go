package model
// 这里的gorm是v1版本
// 数据库连接接口
import (
	"aa/utils"
	"fmt"
	"time"

	_"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func Initdb() {
	// 使用utils里面的数据进行传参来进行路由配置
	// 使用gorm来操作数据库第一个参数是数据库的类型，第二个数据库的配置参数，这里通过fmt包的spriontf参数进行字符串拼接
	db ,err =gorm.Open(utils.DB,fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		utils.DbUser,
		utils.DbPass,
		utils.Dbhost,
		utils.DbPort,
		utils.DbNmae,
	 ))
if err !=nil {
	fmt.Println("数据库连接错误请检查参数")
}
db.SingularTable(true)
// gorm老版本在写入数据库默认会节键值加上复数，调用这个函数参数为true的话会关掉这个功能
// 不使用复数
//创建表,参数是结构体指针
db.AutoMigrate(&User{},&Category{},&Artucle{})

// 迁移数据

db.DB().SetMaxIdleConns(10)
db.DB().SetMaxOpenConns(100)
db.DB().SetConnMaxLifetime(10*time.Second)
	}
	// db, err2 := sql.Open(utils.DB,fmt.Sprintf(
	// 	"%s:%s@tcp(%s:%s)/%s",
	// 	utils.DbUser,
	// 	utils.DbPass,
	// 	utils.Dbhost,
	// 	utils.DbPort,
	// 	utils.DbNmae,
	//  ))
	//  if err2 !=nil {
	// 	fmt.Println("连接数据库失败请检查参数")
	//  }