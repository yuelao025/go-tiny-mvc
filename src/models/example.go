package models

import (
)

type ExampleData struct {
	Message string
	RandomNumber int
}

func GetExampleData() (ExampleData) {

	data := ExampleData {
		Message: "Hello Alice!\n",
		RandomNumber: 10123,
	}
	return data
}