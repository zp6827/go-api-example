package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	// "errors"
)

func testPage(w http.ResponseWriter, r *http.Request) { 
	fmt.Fprintf(w, "Validating IP Address...")
    fmt.Println("Endpoint Hit: validateIpAddress")
}

func handleRequests() { 
	http.HandleFunc("/validateIpAddress", validateIpAddress)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func validateIpAddress(w http.ResponseWriter, r *http.Request) { 
	fmt.Println("Received request at validateIpAddress")
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(reqBody))
}

func main() { 
	handleRequests()
}