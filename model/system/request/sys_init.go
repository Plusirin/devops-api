package request

type InitDB struct {
	Host     string `json:"host"`                        // 服务器地址
	Port     string `json:"port"`                        // 数据库连接端口
	UserName string `json:"username" binding:"required"` // 数据库用户名
	Password string `json:"password"`                    // 数据库密码
	DBName   string `json:"dbname" binding:"required"`   // 数据库名
}
