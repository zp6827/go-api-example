# GoLang API Example

## Running Instructions
* Clone the repository and run `go run main.go` from inside of the `go-api-example` directory. 
* Send POST request to localhost:10000/api/v1/validateIpAddress with the body (json) matching the schema laid out in the openapi.yaml file.

## Future Improvements
* Implement custom error types to better convey to the client why the error occured. 
* Expose the service to gRPC as well. 
    * In order to do this, we would need to separate the server logic from the handling logic. We could have three main functional files: One for defining/handling REST routes, one for defining/handling gRPC routes, and another file for the actual logic of grabbing the country for the IP. 
* In order to escape the concerns of parsing the countries by string name, we could require our service to receive a list of country codes as opposed to country names. We could perform the lookup by country code instead of name (to eliminate errors such as capitalization/misspellings/foreign languages)

## Ideas for Scaling
* We will need to periodically update the database used for mapping IPs to countries. We could create a small service that fetched the DB from the URL and updated the copy in the local filesystem. We could run this on a regular interval (i.e. once daily/weekly). In my searching, I found a package [gocron](https://github.com/go-co-op/gocron) that would allow us to do this very easily.

## Example cURL requests
```
curl -d '{"ipAddress":"74.209.24.0", "validCountries":["United States", "Brazil", "Germany"]}' -H "Content-Type: application/json" -X POST http://localhost:10000/api/v1/validateIpAddress
curl -d '{"ipAddress":"2.16.15.2", "validCountries":["United States", "Brazil", "Germany"]}' -H "Content-Type: application/json" -X POST http://localhost:10000/api/v1/validateIpAddress
curl -d '{"ipAddress":"2.16.9.2", "validCountries":["United States", "Brazil", "Germany"]}' -H "Content-Type: application/json" -X POST http://localhost:10000/api/v1/validateIpAddress
curl -d '{"ipAddress":"1.178.224.1", "validCountries":["United States", "Brazil", "Germany"]}' -H "Content-Type: application/json" -X POST http://localhost:10000/api/v1/validateIpAddress
```
