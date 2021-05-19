# GoLang API Example

## Running Manually
* Clone the repository and run `go run main.go` from inside of the `go-api-example/src` directory. 
* Send POST request to localhost:10000/api/v1/validateIpAddress with the body (json) matching the schema laid out in the openapi.yaml file.
    * **NOTE:** There are sample cURL requests below for testing.
* Exit the program with ctrl+c

## Running Via Docker
* Navigate into `go-api-example` and run `docker build -t go-api-example-dev .`
* Run `docker run --rm -it -p 10000:10000 go-api-example-dev`
* You should now be able to send a POST request to localhost:10000/api/v1/validateIpAddress (see cURL requests or openapi.yml in repository for schema/examples)
* Exit the program with ctrl+c

## Future Improvements
* Implement custom error types to better convey to the client why the error occured. 
* Expose the service to gRPC as well. 
    * In order to do this, we would need to separate the server logic from the handling logic. We could have three main functional files: One for defining/handling REST routes, one for defining/handling gRPC routes, and another file for the actual logic of grabbing the country for the IP. 
* In order to escape the concerns of parsing the countries by string name, we could require our service to receive a list of country codes as opposed to country names. We could perform the lookup by country code instead of name (to eliminate errors such as capitalization/misspellings/foreign languages)
* Right now, I believe that the methods called from `handleValidateIpAddress` probably do too much. Its throwing red flags that we are possibly returning three different errors from one single function. If I had more time, I would break these functions out to simplify the error-handling and general readability. 

## Ideas for Scaling
* We will need to periodically update the database used for mapping IPs to countries. We could create a small service that fetched the DB from the URL and updated the copy in the local filesystem. We could run this on a regular interval (i.e. once daily/weekly). In my searching, I found a package [gocron](https://github.com/go-co-op/gocron) that would allow us to do this very easily.
* We could separate a lot of this logic out to their own files as opposed to having everything exist in main.go. For example, custom types and general logic can move out to other files as we expand. 
* Implement a better approach to versioning. Right now, all endpoints are to api/v1/validateIpAddress. As we improve this API and possibly add more endpoints, we can branch off by version number and route requests based on the API version. I.e. api/v1, api/v2. . .
* Allow the IP/Port to be set via environment variables as opposed to being hardcoded into the files

## Example cURL requests
```bash
# USA IP - valid
curl -d '{"ipAddress":"74.209.24.0", "validCountries":["United States", "Brazil", "Germany"]}' -H "Content-Type: application/json" -X POST http://localhost:10000/api/v1/validateIpAddress
# Brazil IP - valid
curl -d '{"ipAddress":"2.16.15.2", "validCountries":["United States", "Brazil", "Germany"]}' -H "Content-Type: application/json" -X POST http://localhost:10000/api/v1/validateIpAddress
# Germany IP - valid
curl -d '{"ipAddress":"2.16.9.2", "validCountries":["United States", "Brazil", "Germany"]}' -H "Content-Type: application/json" -X POST http://localhost:10000/api/v1/validateIpAddress
# Spain IP - not valid
curl -d '{"ipAddress":"1.178.224.1", "validCountries":["United States", "Brazil", "Germany"]}' -H "Content-Type: application/json" -X POST http://localhost:10000/api/v1/validateIpAddress
```
