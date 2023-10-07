package main

import (
	"demo/data"
	"demo/router"
)

func main() {
	data.InitDb()
	router.InitRouter()
	
}
