package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB
func init(){
	var err error
	//sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/db_movie")
	Db, err = gorm.Open("mysql", "root:8084810821@tcp(106.13.107.166:3306)/db_movie")
	Db.SingularTable(true)
	Db.LogMode(true)
	if err != nil{
		panic(err)
	}
}