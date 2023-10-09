package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

var d *gorm.DB


func Connect(){
	username := "root"
	password := "ayushsql"
	host := "localhost"
	port := "3306"
	dbname := "bookstore"
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}
	d=db

}

func GetDB() *gorm.DB{
	return d
}


