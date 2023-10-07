package api

import (
	"demo/data"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetInfo(c *gin.Context) {

	info, total := data.GetAllInfo()
	c.JSON(http.StatusOK, gin.H{
		"data": info,
		"num":  total,
	})

}
func GetInfoByName(c *gin.Context) {

	city := c.Query("name")

	info := data.GetInfoByCity(city)
	c.JSON(http.StatusOK, gin.H{
		"data": info,
	})

}

func GetInfoById(c *gin.Context) {
	left, _ := strconv.Atoi(c.PostForm("left"))
	info := data.GetInfosById(left)
	c.JSON(http.StatusOK, gin.H{
		"data": info,
	})
}
