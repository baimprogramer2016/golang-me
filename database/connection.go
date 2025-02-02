package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMysql() (*gorm.DB, error) {
	dsn := "root:P@ssw0rd@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
