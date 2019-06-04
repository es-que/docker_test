package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const API_URL = "https://api.macaddress.io/v1"

var ERR_NONE_200_RESPONSE = errors.New("Got a none 200 response")

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	mac_address := os.Args[1]
	if mac_address == "-h" || mac_address == "help" || mac_address == "--help" {
		usage()
	}

	api_key := os.Getenv("API_KEY")
	err := execRequest(mac_address, api_key)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

func execRequest(mac_address, api_key string) error {
	url := fmt.Sprintf("%s?output=vendor&search=%s", API_URL, mac_address)
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Failed to create new request, %s\n", err.Error())
		return err
	}
	request.Header.Set("X-Authentication-Token", api_key)

	response, err := client.Do(request)
	if err != nil {
		log.Printf("Failed to complete API call to macaddress.io, %s\n", err.Error())
		return err
	}
	defer response.Body.Close()

	err = checkResponseForError(response.StatusCode)
	if err != nil {
		return err
	}

	company_name := bytes.NewBuffer(nil)
	_, err = io.Copy(company_name, response.Body)
	if err != nil {
		log.Printf("Failed to read response, %s\n", err.Error())
		return err
	}

	fmt.Printf("%s\n", string(company_name.Bytes()))
	return nil
}

func checkResponseForError(status_code int) error {
	var message = ""
	var err error
	switch status_code {
	case 200:
		err = nil
	case 400:
		message = "Invalid parameters."
		err = ERR_NONE_200_RESPONSE
	case 401:
		message = "Access restricted. Enter the correct API key."
		err = ERR_NONE_200_RESPONSE
	case 402:
		message = "Access restricted. Check the credits balance."
		err = ERR_NONE_200_RESPONSE
	case 422:
		message = "Invalid MAC or OUI address was received."
		err = ERR_NONE_200_RESPONSE
	case 429:
		message = "Too many requests. Try your call again later."
		err = ERR_NONE_200_RESPONSE
	case 500:
		message = "Internal server error. Try your call again or contact us."
		err = ERR_NONE_200_RESPONSE
	}
	if err != nil {
		log.Printf("%s - %d, %s", err.Error(), status_code, message)
	}
	return err
}

func usage() {
	fmt.Printf("%s - <mac address>\n", os.Args[0])
	os.Exit(1)
}
