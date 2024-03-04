package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:04271017@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=true"

func NewConn() (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
