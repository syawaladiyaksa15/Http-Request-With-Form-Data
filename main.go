package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var baseUrl = "http://localhost:8001"

type mahasiswa struct {
	NPM   string
	Nama  string
	Grade int
}

func fetchUser(NPM string) (mahasiswa, error) {
	var err error
	var client = &http.Client{}
	var data mahasiswa

	var param = url.Values{}
	param.Set("npm", NPM)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseUrl+"/user", payload)
	if err != nil {
		return data, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)

	if err != nil {
		return data, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func main() {
	var user, err = fetchUser("2015030031")
	if err != nil {
		fmt.Println("Error!!!", err.Error())
		return
	}

	fmt.Printf("NPM : %s, Nama : %s, Grade : %d\n", user.NPM, user.Nama, user.Grade)

}
