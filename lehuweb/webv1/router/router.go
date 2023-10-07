package router

import (
	"github.com/gin-gonic/gin"
	"webv1/api"
	"webv1/midware"
	"webv1/utils"
)

type form struct {
	name string
}

func InitRouter() {
	gin.SetMode(utils.AppMode)
	router := gin.Default()
	router.Use(midware.Cors())
	//router.GET("/", func(c *gin.Context) {
	//	c.ShouldBind(&form{})
	//})
	router.GET("/api", api.GetInfo)
	_ = router.Run(utils.HttpPort)

}

//curl  POST -d '{"left": "10", "right":30}' "http://localhost:3000/id"

//curl -v --form left=10 --form right=20 http://localhost:3000/id
