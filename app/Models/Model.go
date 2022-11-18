package Models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
}

func init() {
	//我这里用到数据库是mysql，需要配置DSN属性[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	dsn := "root:123456@tcp(127.0.0.1:11002)/go_test?charset=utf8&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}

}