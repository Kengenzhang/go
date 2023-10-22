package utils
// 配置文件数据接口
import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	DB       string
	Dbhost   string
	DbPort   string
	DbUser   string
	DbPass   string
	DbNmae   string
)
// 接收配置数据的容器
// 数据初始化
func init() {
	// ini/load 是init包中的一个方法用于获取配置文件的数据，返回一个file的指针和合格err
	f, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误请检查文件路径")
	}
	LoadServer(f)
	// 载入配置服务的数据
	Loadata(f)
	// 从ini文件中获取数据库配置信息的数据

}
func LoadServer(f *ini.File) {
	/*这里就是通过f这个指针去调用section这个方法方法里面写ini文件的区段，
	key这个就是去找区段里面的键值，如果没有返回的话通过函数must string函数返回默认值

	*/
	AppMode = f.Section("server").Key("AppMode").MustString("debug")
	HttpPort = f.Section("server").Key("HttpPort").MustString(":3000")
}
func Loadata(f *ini.File) {
	// 函数section获取ini文件的节，key函数获取参数里面的值，muststing函数是设置的默认值
	DB = f.Section("database").Key("DB").MustString("mysql")
	Dbhost = f.Section("database").Key("Dbhost").MustString("localhost")
	DbNmae = f.Section("database").Key("Dbname").MustString("ginname")
	DbPort = f.Section("database").Key("DbPort").MustString("3306")
	DbUser = f.Section("database").Key("DbUser").MustString("ginblog")
	DbPass = f.Section("database").Key("DbPass").MustString("admin123")
}
//将config里面的数据传过来