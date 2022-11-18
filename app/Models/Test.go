package Models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  uint
}

func init() {
	//我这里用到数据库是mysql，需要配置DSN属性[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	dsn := "root:123456@tcp(127.0.0.1:11002)/go_test?charset=utf8&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}

	// 迁移 schema
	db.AutoMigrate(&User{})

	// Insert 插入语句
	db.Create(&User{Name: "linzy", Age: 23})

	// Select 查询语句
	var user User
	db.First(&user, 1) // 根据整形主键查找
	fmt.Println(user)
	db.First(&user, "name = ?", "linzy") // 查找 name 字段值为 linzy 的记录
	fmt.Println(user)

	// Update 更新语句 - 将 User 的 age 更新为 18
	db.Model(&user).Update("Age", 18)
	// Update - 更新多个字段
	db.Model(&user).Updates(User{Name: "linzy", Age: 88}) // 仅更新非零值字段
	db.Model(&user).Updates(map[string]interface{}{"Name": "linzy", "Age": 23})

	// Delete 删除语句 - 删除 ID为1
	db.Delete(&user, 1)
}