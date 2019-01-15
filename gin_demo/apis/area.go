package apis

import (
	"github.com/gin-gonic/gin"
	"log"
	"lovehome/gin_demo/models"
	"net/http"
)

func GetArea(c *gin.Context) {
	p := models.Area{}

	areas, err := p.GetArea()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"areas": areas,
	})
}
