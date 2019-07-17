package main

import (
	"fmt"
	"./secret"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func gormConnect() *gorm.DB {
	auth_data := secret.GetAuthData()

	CONNECT := auth_data[1] + ":" + auth_data[2] + "@" + auth_data[3] + "/" + auth_data[4]
	// DBに接続
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
	// 最後にDB接続を切断
	defer db.Close()

	langFirst := lang{}
	// SELECT * FROM langs ORDER BY id LIMIT 1;
	db.First(&langFirst)
	fmt.Println(langFirst)

	langTake := lang{}
	// SELECT * FROM langs LIMIT 1;
	db.Take(&langTake)
	fmt.Println(langTake)

	// SELECT * FROM langs ORDER BY id DESC LIMIT 1;
	langLast := lang{}
	db.Last(&langLast)
	fmt.Println(langLast)

	// SELECT * FROM langs;
	langAll := []lang{}
	db.Find(&langAll)
	fmt.Println(langAll)

	// SELECT * FROM langs WHERE id = 3;
	langSome := lang{}
	db.First(&langSome, 4)
	fmt.Println(langSome)

	// SELECT * FROM langs WHERE name = "PHP" LIMIT 1;
	langWhereFirst := lang{}
	db.Where("name = ?", "PHP").First(&langWhereFirst)
	fmt.Println(langWhereFirst)

	// SELECT * FROM langs WHERE name = 'Java';
	langWhereFind := []lang{}
	db.Where("name = ?", "Java").Find(&langWhereFind)
	fmt.Println(langWhereFind)

	// SELECT * FROM langs WHERE name <> 'Ruby'
	langWhereNotFind := []lang{}
	db.Where("name <> ?", "Ruby").Find(&langWhereNotFind)
	fmt.Println(langWhereNotFind)

	// SELECT * FROM langs WHERE name IN ('PHP', 'Ruby');
	langWhereIn := []lang{}
	db.Where("name IN (?)", []string{"PHP", "Ruby"}).Find(&langWhereIn)
	fmt.Println(langWhereIn)
}
