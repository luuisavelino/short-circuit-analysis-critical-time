package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/short-circuit-analysis-critical-time/models"
	"github.com/luuisavelino/short-circuit-analysis-critical-time/pkg/analysis"
)

func GetBarraCurtoCircuito(c *gin.Context) string {
	line := c.Params.ByName("line")
	point := c.Params.ByName("point")

	parts := strings.Split(line, "-")

	if point == "0" {
		return parts[0]
	} else if point == "100" {
		return parts[1]
	}

	return "ficticia"
}

func jsonError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"Error": err.Error(),
	})
}

func GetAnalysis(BarraCurtoCircuito string) map[string]map[string]float64 {
	var systemInfo = make(map[string]map[string]float64)

	for gerador, data := range models.Elements["1"] {
		posicao_barra_gerador_pre := models.BarrasAdicionadasBefore[gerador].Posicao
		posicao_barra_cc_pre := models.BarrasAdicionadasBefore[BarraCurtoCircuito].Posicao

		posicao_barra_gerador_pos := models.BarrasAdicionadasAfter[gerador].Posicao
		posicao_barra_cc_pos := models.BarrasAdicionadasAfter[BarraCurtoCircuito].Posicao

		X_1_2_pre := models.AllZbusBeforeFault["positiva"][posicao_barra_gerador_pre][posicao_barra_cc_pre]
		X_1_2_pos := models.AllZbusAfterFault["positiva"][posicao_barra_gerador_pos][posicao_barra_cc_pos]
		Xd := data.Z_positiva

		systemInfo[gerador] = analysis.SystemAnalysis(X_1_2_pre, X_1_2_pos, Xd)
	}

	return systemInfo
}

func Data(c *gin.Context) {
	fileId := c.Params.ByName("fileId")
	line := c.Params.ByName("line")
	point := c.Params.ByName("point")

	var ch1, ch2, ch3, ch4, ch5 = make(chan []byte), make(chan []byte), make(chan []byte), make(chan []byte), make(chan []byte)


	var err = make(chan error)

	var wg sync.WaitGroup
	wg.Add(5)

	go getDataFromAPI(&wg, ch1, err, host+":"+port+"/api/v2/files/"+fileId+"/zbus/short-circuit/"+line+"/point/"+point)
	json.Unmarshal(<-ch1, &models.AllZbusBeforeFault)

	go getDataFromAPI(&wg, ch2, err, host+":"+port+"/api/v2/files/"+fileId+"/zbus/atuacao/"+line)
	json.Unmarshal(<-ch2, &models.AllZbusAfterFault)

	go getDataFromAPI(&wg, ch3, err, host+":"+"8080"+"/api/v2/files/"+fileId+"/types/1/elements")
	json.Unmarshal(<-ch3, &models.Elements)

	go getDataFromAPI(&wg, ch4, err, host+":"+port+"/api/v2/files/"+fileId+"/zbus/short-circuit/"+line+"/point/"+point+"/bars")
	json.Unmarshal(<-ch4, &models.BarrasAdicionadasBefore)

	go getDataFromAPI(&wg, ch5, err, host+":"+port+"/api/v2/files/"+fileId+"/zbus/atuacao/"+line+"/bars")
	json.Unmarshal(<-ch5, &models.BarrasAdicionadasAfter)
	
	if <-err != nil {
		fmt.Println("Erro ao obter dados da API:", <-err)
		jsonError(c, <-err)
		return
	}

	wg.Wait()
}

func getDataFromAPI(wg *sync.WaitGroup, c chan<- []byte, e chan<- error, url string) {
	wg.Done()
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("erro 1")
		e <- err
		return
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("erro 2")
		e <- err
		return
	}

	c <- responseData
	e <- err
}
