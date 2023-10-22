package model

// 用户模块
import (
	"aa/utils/errmsg"
	"encoding/base64"
	"log"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	// "github.com/ugorji/go/codec"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// 查询用户有没有
func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username=?", name).First(&users)
	if users.ID > 0 {
		// id大于0说明用户存在
		return errmsg.ERROR_USERNAME //1001
	}
	return errmsg.SUCCSE
}

// 新增用户
func CreateUser(data *User) int {
	data.Password = Scryppw(data.Password)
	err2 := db.Create(&data).Error
	if err2 != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE
}

// 查询用户列表
func Getusers(pagesize int, pagenum int) []User {
	var users []User
	err2 := db.Limit(pagesize).Offset((pagenum - 1) * pagesize).Find(&users).Error
	if err2 != nil && err2 != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// 删除用户
func Deluser(id int) int {
	var user User
	err2 := db.Where("id=?", id).Delete(&user).Error
	if err2 != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
// 编辑用户
func Ususer(id int, data *User) int {
	var user User
	// 不能修改密码，限定在密码之外的
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err2 := db.Model(&user).Where("id=?", id).Update(maps).Error
	if err2 != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
	// 从user{}这个模型传过去，位置是id，跟新的数据是maps
	//创建一个
}

// 密码加密
func Scryppw(password string) string {
	const Keyleng = 10
	salt := make([]byte, 8)
	salt = []byte{66, 89, 77, 12, 45, 52, 30, 19}                            //盐
	Hashpw, err2 := scrypt.Key([]byte(password), salt, 16384, 8, 1, Keyleng) //进行加盐
	if err2 != nil {
		log.Fatal(err2)
	}
	fpw := base64.StdEncoding.EncodeToString(Hashpw) //进行编码
	return fpw                                       //返回加密后的密码

}
