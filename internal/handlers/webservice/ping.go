package webservice

import "github.com/gin-gonic/gin"

func (w *Webservice) GetPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
