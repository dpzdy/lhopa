package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webv1/data"
)

func GetInfo(c *gin.Context) {

	info := data.GetAllInfo()
	c.JSON(http.StatusOK, gin.H{
		"data": info,
	})

}
