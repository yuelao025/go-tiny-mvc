package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"go-web-framework/models"
)


func Sku(w http.ResponseWriter, r *http.Request) {

	// Get data from model
	modelData := models.SkuInfo()

	rlt,err := json.Marshal(modelData)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(rlt))

	// 这个只能是控制台上
	//os.Stdout.Write(rlt)

}



func UserInfo(w http.ResponseWriter, r *http.Request) {

	// Get data from model
	modelData := models.UserInfo()

	rlt,err := json.Marshal(modelData)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(rlt))
	// 这个只能是控制台上
	//os.Stdout.Write(rlt)

}