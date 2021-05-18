package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	// "io/ioutil"
	// "errors"
)

type RequestBody struct { 
	IpAddress string
	ValidCountries []string
}

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
	decoder := json.NewDecoder(r.Body)
    var requestBody RequestBody
    err := decoder.Decode(&requestBody)
    if err != nil {
        panic(err)
    }

    fmt.Println("IP ADDRESS:", requestBody.IpAddress)
    fmt.Println("Valid Countries:", requestBody.ValidCountries)	
}

func main() { 
	handleRequests()
}