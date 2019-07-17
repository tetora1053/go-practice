package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func gormConnect() *gorm.DB {
	auth_data := GetAuthData()

	CONNECT := auth_data[1] + ":" + auth_data[2] + "@" + auth_data[3] + "/" + auth_data[4]
	db, err := gorm.Open(auth_data[0], CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type lang struct {
	Id int `json:id`
	Name string `json:name`
}

func main() {
	db := gormConnect()
	defer db.Close()

	lang := lang{}
	lang.Id = 2
	db.First(&lang)
	fmt.Println(lang)
}
