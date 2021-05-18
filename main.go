package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/IncSW/geoip2"
	"net"
	// "io/ioutil"
	// "errors"
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
	reader, err := geoip2.NewCountryReaderFromFile(dbPath)
	if err != nil { 
		panic(err)
	}
	countryReader = reader
	fmt.Println("Initialized Country Reader")
}

func handleRequests() { 
	http.HandleFunc("/validateIpAddress", handleValidateIpAddress)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func handleValidateIpAddress(w http.ResponseWriter, r *http.Request) { 
	fmt.Println("Received request at validateIpAddress")
	decoder := json.NewDecoder(r.Body)
    var requestBody RequestBody
    err := decoder.Decode(&requestBody)
    if err != nil {
        panic(err)
    }

	countryName := getCountryNameForIpAddress(requestBody.IpAddress)
	isCountryValid := contains(requestBody.ValidCountries, countryName)

	if isCountryValid { 
		fmt.Println("Country Name is Valid")
	} 

	response := Response{isCountryValid, ""}
	responseJson, err := json.Marshal(response)
	if err != nil { 
		panic(err)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
	// fmt.Println("Response", response)
}

func getCountryNameForIpAddress(ipAddress string) string { 
	parsedIp := net.ParseIP(ipAddress)
	if parsedIp == nil {
		panic("err")
	}
	record, err := countryReader.Lookup(parsedIp)
	if err != nil {
		panic(err)
	}
	countryName := record.Country.Names["en"]
	return countryName
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