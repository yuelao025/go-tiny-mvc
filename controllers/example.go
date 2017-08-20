package controllers

import (
	//"github.com/sirupsen/logrus"
	"net/http"
	"encoding/json"
	// For API only services it might
	// not be neccessary to use views.
	//"go-web-framework/views"


	// Use models to access data.
	"go-web-framework/models"
	"fmt"
)

func Example(w http.ResponseWriter, r *http.Request) {


	/*
		Example web page.
	*/

	//// Logging the start and end of the request.
	//logrus.WithFields(logrus.Fields{
	//	"stage":    "controller",
	//	"function": "Example()",
	//	"url":      r.URL,
	//	"resp":		w,
	//}).Debug("Request recieved")

	//  请求结束 表示
	//defer logrus.WithFields(logrus.Fields{
	//	"stage":    "controller",
	//	"function": "Example()",
	//	"url":      r.URL.Path,
	//}).Debug("Request processed")

	// Get data from model
	modelData := models.GetExampleData()

	rlt,err := json.Marshal(modelData)

	if err != nil {
		fmt.Println(err)
	}

	// 将 json 输出
	fmt.Fprintf(w, string(rlt))
	// 这个只能是控制台上
	//os.Stdout.Write(rlt)

	//fmt.Println(rlt)

// 将 view x 掉；
	//// Transform this into data for view
	//viewData := views.ExampleData {
	//	Message: modelData.Message,
	//}
	//
	//// Pass the data to the view.
	//views.Example(w, viewData)

}
