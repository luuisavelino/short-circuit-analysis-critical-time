package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var (
	val1 []byte
	val2 []byte
	val3 []byte
)

func main() {
	var wg sync.WaitGroup
	erro := make(chan error)

	wg.Add(3)

	go getDataAPI1(&wg, erro)
	go getDataAPI2(&wg, erro)
	go getDataAPI3(&wg, erro)

	if <-erro != nil {
		return
	}

	go func() {
		wg.Wait()
		close(erro)
	}()


	json.Unmarshal(val1, &resp1)
	json.Unmarshal(val2, &resp2)
	json.Unmarshal(val3, &resp3)

	fmt.Println("Valor 1:", resp1)
	fmt.Println("Valor 2:", resp2)
	fmt.Println("Valor 3:", resp3)
}

func getDataAPI1(wg *sync.WaitGroup, erro chan error) {
	defer wg.Done()
	data, err := getDataFromAPI("http://localhost:8080/api/v2/files/1/size")
	if err != nil {
		fmt.Println("Erro ao obter dados da API 1:", err)
	}
	val1 = data
	erro <- err
}

func getDataAPI2(wg *sync.WaitGroup, erro chan error) {
	defer wg.Done()
	data, err := getDataFromAPI("http://localhost:8083/api/v2/files/1/size")
	if err != nil {
		fmt.Println("Erro ao obter dados da API 2:", err)
	}
	val2 = data
	erro <- err
}

func getDataAPI3(wg *sync.WaitGroup, erro chan error) {
	defer wg.Done()
	data, err := getDataFromAPI("http://localhost:8080/api/v2/files/1/size")
	if err != nil {
		fmt.Println("Erro ao obter dados da API 3:", err)
	}
	val3 = data
	erro <- err
}

func getDataFromAPI(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

var resp1 = make(map[string]int)
var resp2 = make(map[string]int)
var resp3 = make(map[string]int)
