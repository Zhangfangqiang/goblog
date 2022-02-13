package model

import (
	"goblog/pkg/logger"
	"gorm.io/driver/mysql" //GORM 的 MySQL 数据库驱动导入
	"gorm.io/gorm"
)

var DB *gorm.DB //DB gorm.DB 对象, 相当于静态访问

/**
 * 连接数据库 初始化模型
 */
func ConnectDB() *gorm.DB {

	var err error
	config := mysql.New(mysql.Config{
		DSN: "root:123456@tcp(127.0.0.1:3306)/goblog?charset=utf8&parseTime=True&loc=Local",
	})

	// 准备数据库连接池
	DB, err = gorm.Open(config, &gorm.Config{})

	logger.LogError(err)

	return DB
}
