package router

import (
	"demo/api"
	"demo/midware"
	"demo/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	router := gin.Default()
	router.Use(midware.Cors())
	router.GET("/api", api.GetInfo)
	router.GET("/item", api.GetInfoByName)
	router.POST("/id", api.GetInfoById)
	_ = router.Run(utils.HttpPort)

}

//curl  POST -d '{"left": "10", "right":30}' "http://localhost:3000/id"

//curl -v --form left=10 --form right=20 http://localhost:3000/id
