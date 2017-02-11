package db

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "github.com/alexandercrosson/gingorm/models"
)

var db *gorm.DB
var err error 

func Init() {
    db, _ = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/gingorm?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        fmt.Println(err)
    }

    
    db.AutoMigrate(&models.Person{})
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
    db.Close()
}
