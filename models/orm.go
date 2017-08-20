package models

import (
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"github.com/jinzhu/gorm"
)

//func init()  {
//
//	sql.Register("mysql",mysql.MySQLDriver{})
//
//}

func GetDb() (* gorm.DB)  {


	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/blog2?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("failed to open database:", err.Error())
		return nil
	}

	return  db


	//defer db.Close()



	//db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/blog2?charset=utf8")
	//if err != nil {
	//	fmt.Println("failed to open database:", err.Error())
	//	return nil
	//}
	////defer db.Close()
	//return  db
}