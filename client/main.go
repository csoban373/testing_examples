package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	MakeGet("/ping")
}

func MakeGet(target string) (string, error) {
	client := http.Client{}
	response, err := client.Get(ApiEndpoint + target)
	// do stuff with response...
	if err != nil || response.StatusCode != 200 {
		fmt.Println("client get failed!")
		return "", err
	}
	defer response.Body.Close()

	type Response struct {
		Message string `json:"message"`
	}
	var decodedResponse Response
	err = json.NewDecoder(response.Body).Decode(&decodedResponse)
	if err != nil {
		fmt.Println("decode failed!")
		return "", err
	}

	fmt.Println(decodedResponse.Message)

	return decodedResponse.Message, nil
}
