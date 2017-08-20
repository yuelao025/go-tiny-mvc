package models

import (

)

type Sku struct {
	Goods_SN string
	Name string
}

type Tbl_User struct {
	Username string
	Password string
	Email string
}

func SkuInfo() (Sku)  {
	
	data := Sku{
		Goods_SN:"skuno11111",
		Name:"iphone6",
	}

	return data

}

func UserInfo() ([]Tbl_User) {

	db := GetDb()

	user := []Tbl_User{}

	db.Find(&user)

	return user

	//if db == nil{
	//	return []User{}
	//}
	//rows,err := db.Query("SELECT  * from tbl_user")
	//if err != nil{
	//	log.Println("query failed ! %v",err)
	//	return []User{}
	//}
	//
	//defer rows.Close()
	//
	//
	//if rows.Next(){
	//	rows.Scan(&users)
	//}
	//
	//log.Println("%v",users)
	//
	//return users

}

