package errmsg

// 错误信息
const (
	SUCCSE = 200
	ERROR  = 500
	// code =1000 用户模块错误
	ERROR_USERNAME     = 1001 //用户名被使用
	ERROR_PASSWORD     = 1002 //密码错误
	ERROR_USER_NOT     = 1003 //用户不存在
	ERROR_TOKEN_NOT    = 1004 //用户的输入的token不存在
	ERROR_TOKEN_NOTIME = 1005 //用户token超时
	ERROR_TOKEN_WRONG  = 1006 //TOKEN wrong
	ERROR_TOKEN_TYPE   = 1007 //TOKEN格式不对

		//code 3000 分类模块错误
	ERROR_CATEGORY=2001 //查询分类
	// code=2000 文章模块错误
	ERROR_ARTNOT=3001
	ERROR_CATE_NOT_EXIST=3002
)

// 返回错误信息列表
var Codemsg = map[int]string{
	SUCCSE:             "OK",
	ERROR:              "FATL",
	ERROR_USERNAME:     "用户已存在",
	ERROR_PASSWORD:     "密码错误",
	ERROR_USER_NOT:     "用户不存在",
	ERROR_TOKEN_NOT:    "TOKEN不存在",
	ERROR_TOKEN_NOTIME: "TOKEN超时",
	ERROR_TOKEN_WRONG:  "TOKEN错误",
	ERROR_TOKEN_TYPE:   "TOKEN格式不对",
	//分类
	ERROR_CATEGORY:"该分类已存在",
	ERROR_ARTNOT:"文章不存在",
	ERROR_CATE_NOT_EXIST:"分类不存在,查询不到文章",
}

// 返回错误信息
// 通过接收状态码，然后返回错误信息
func Geterrmsg(code int) string {
	return Codemsg[code]
}
