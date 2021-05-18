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

var countryReader *geoip2.CountryReader

func init() { 
	reader, err := geoip2.NewCountryReaderFromFile("testdata/GeoLite2-Country.mmdb")
	if err != nil { 
		panic(err)
	}
	countryReader = reader
	fmt.Println("Initialized Country Reader")
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

	countryName := getCountryNameForIpAddress(requestBody.IpAddress)

    fmt.Println("IP ADDRESS:", requestBody.IpAddress)
    fmt.Println("Valid Countries:", requestBody.ValidCountries)	
	fmt.Println("COUNTRY:", countryName)
}

func getCountryNameForIpAddress(ipAddress string) string { 
	record, err := countryReader.Lookup(net.ParseIP(ipAddress))
	countryName := record.Country.Names["en"]
	if err != nil {
		panic(err)
	}
	return countryName
}

func main() { 
	handleRequests()
}