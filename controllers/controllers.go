package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var host string = "http://localhost" //os.Getenv("elements_host")
var port string = "8081"             //os.Getenv("elements_port")

func AllData(c *gin.Context) {
	Data(c)

	barraCurtoCircuito := GetBarraCurtoCircuito(c)

	SystemInfo := GetAnalysis(barraCurtoCircuito)

	c.JSON(http.StatusOK, SystemInfo)
}

func GeradorData(c *gin.Context) {
	gerador := c.Params.ByName("gerador")

	Data(c)

	barraCurtoCircuito := GetBarraCurtoCircuito(c)

	SystemInfo := GetAnalysis(barraCurtoCircuito)

	c.JSON(http.StatusOK, SystemInfo[gerador])
}
