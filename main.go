package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Structs should represent the JSON string
type astronauts struct {
	People  []people `json:"people"`
	Number  int      `json:"number"`
	Message string   `json:"message"`
}

type people struct {
	Name  string `json:"name"`
	Craft string `json:"craft"`
}

func main() {
	// Declare the location of the JSON string
	url := "http://api.open-notify.org/astros.json"

	// Create an Http Client and an Http Request
	spaceClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set proper headers
	req.Header.Set("User-Agent", "go-2-space")

	// Send the request, and read the body of the response
	res, resErr := spaceClient.Do(req)
	if resErr != nil {
		log.Fatal(resErr)
	}

	body, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		log.Fatal(bodyErr)
	}

	// Unmarshal the body into the struct
	astros := astronauts{}
	json.Unmarshal(body, &astros)
	fmt.Println(astros.People[0].Name)
}
