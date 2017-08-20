package models

import (
)

type ExampleData struct {
	Message string
	RandomNumber int
}



func GetExampleData() (ExampleData) {

	data := ExampleData {
		Message: "default  value !\n",
		RandomNumber: 10123,
	}
	return data
}

