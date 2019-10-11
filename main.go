package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type astronaughts struct {
	People []people `json:"people"`
	Number int      `json:"number"`
	//name string `json:"name"` // type: name of type
	//craft string `json:"craft"`
	Message string `json:"message"`
}

type people struct {
	Name  string `json:"name"`
	Craft string `json:"craft"`
}

func main() {
	url := "http://api.open-notify.org/astros.json"
	spaceClient := http.Client{
		Timeout: time.Second * 5, //5 sec to get data.
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "go-2-space")

	res, resErr := spaceClient.Do(req)
	if resErr != nil {
		log.Fatal(resErr)
	}
	body, bodyErr := (ioutil.ReadAll(res.Body))
	if bodyErr != nil {
		log.Fatal(bodyErr)
	}

	astros := astronaughts{}
	json.Unmarshal(body, &astros)
	fmt.Println(astros.Number)
	fmt.Println(astros.Message)
	fmt.Println(astros)
	//fmt.Println(res.Body)
}
