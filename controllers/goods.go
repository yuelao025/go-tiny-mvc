package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"go-web-framework/models"
	"os"
)


func Sku(w http.ResponseWriter, r *http.Request) {

	// Get data from model
	modelData := models.SkuInfo()

	rlt,err := json.Marshal(modelData)

	if err != nil {
		fmt.Println(err)
	}

	os.Stdout.Write(rlt)

}
