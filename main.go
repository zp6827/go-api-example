package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/IncSW/geoip2"
	"net"
	"errors"
)

type RequestBody struct { 
	IpAddress string
	ValidCountries []string
}

type Response struct { 
	IsCountryValid bool `json:"isCountryValid"`
 	ErrorString string `json:"errorString"`
}

var countryReader *geoip2.CountryReader
const dbPath = "testdata/GeoLite2-Country.mmdb"

func init() { 
	// Instantiate the db reader on init
	reader, err := geoip2.NewCountryReaderFromFile(dbPath)
	if err != nil { 
		panic(err)
	}
	countryReader = reader
	fmt.Println("Initialized Country Reader")
}

func handleRequests() { 
	http.HandleFunc("/api/v1/validateIpAddress", handleValidateIpAddress)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

// Handle the request and respond appropriately
func handleValidateIpAddress(w http.ResponseWriter, r *http.Request) { 
	fmt.Println("Received request at validateIpAddress")
	decoder := json.NewDecoder(r.Body)
    var requestBody RequestBody
    err := decoder.Decode(&requestBody)
    if err != nil {
		errorString := `Failed to decode request body. 
		Ensure request is valid JSON and contains fields "ipAddress" and "validCountries"`
        writeResponse(w, false, errorString, 400)
		return
    }

	countryName, err := getCountryNameForIpAddress(requestBody.IpAddress)
	if err != nil {
		writeResponse(w, false, err.Error(), 400)
		return
	}

	isCountryValid := contains(requestBody.ValidCountries, countryName)

	writeResponse(w, isCountryValid, "", 200)
	return
}

// Wrapper function for creating/writing the response
func writeResponse(w http.ResponseWriter, isCountryValid bool, errorString string, httpStatusCode int) { 
	w.Header().Set("Content-Type","application/json")

	response := Response{isCountryValid, errorString}
	responseJson, err := json.Marshal(response)

	// TODO: how do we handle this
	if err != nil { 
		fmt.Println("OH NO, FAILED TRYING TO CONVERT THIS TO JSON")
	}

	w.WriteHeader(httpStatusCode)
	w.Write(responseJson)
}

// Returns the country name associated with the given IP address
func getCountryNameForIpAddress(ipAddress string) (string, error) { 
	parsedIp := net.ParseIP(ipAddress)
	if parsedIp == nil {
		err := errors.New("Unable to parse IP Address")
		return "", err
	}

	record, err := countryReader.Lookup(parsedIp)
	if err != nil {
		return "", err
	}

	countryName := record.Country.Names["en"]
	return countryName, nil
}

// Utility function to return whether a value exists in a slice
func contains(s []string, str string) bool { 
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func main() { 
	handleRequests()
}