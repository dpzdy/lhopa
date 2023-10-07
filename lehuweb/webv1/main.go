package main

import (
	"webv1/data"
	"webv1/router"
)

func main() {
	data.InitDb()
	router.InitRouter()
	//data.GetAllInfo()

}
