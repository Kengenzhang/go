package main

import (
	"aa/model"
	router "aa/routes"
)

func main() {
	model.Initdb()
	router.InitRouter()

}