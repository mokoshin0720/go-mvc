// gormに依存するstruct
package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	// gorm.Model // ID(プライマリキー), CreatedAt, UpdatedAt, DeletedAtをもつ
	Name string `json:"name"`
	Id uint `json:"id"`
}

type Sample struct {
	Message string `json:"message"`
}

func GetDBConn() *gorm.DB {
	dsn := "root:password@tcp(db:3306)/dbmate"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	   panic(err)
	}
	return db
}