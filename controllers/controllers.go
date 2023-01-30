package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var host string = "http://localhost" //os.Getenv("elements_host")
var port string = "8081"             //os.Getenv("elements_port")

func AllData(c *gin.Context) {
	err := Data(c)
	if err != nil {
		jsonError(c, err)
		return
	}

	barraCurtoCircuito := GetBarraCurtoCircuito(c)

	SystemInfo := GetAnalysis(barraCurtoCircuito)

	c.JSON(http.StatusOK, SystemInfo)
}

func GeradorData(c *gin.Context) {
	gerador := c.Params.ByName("gerador")

	err := Data(c)
	if err != nil {
		jsonError(c, err)
		return
	}

	barraCurtoCircuito := GetBarraCurtoCircuito(c)

	SystemInfo := GetAnalysis(barraCurtoCircuito)

	c.JSON(http.StatusOK, SystemInfo[gerador])
}
