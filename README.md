# NWS-API-Weather-Forecast
Creates an HTTP Server that tells the client what the weather is based on NWS data

## Set up

1. Check if Go Cli is installed on your machine (run `go version`)
2. If not, then download version 1.24 of Go: https://go.dev/doc/install

## Run Server

To run the server, run in your terminal:

```
cd nws-api
go clean
go build
go run main.go
```

Once you see the server is up (by default it run on port 8080), make a POST call to: `http://localhost:8080/today-forecast`

Here's an example of how the cURL command should look like:

```
curl --location 'http://localhost:8080/today-forecast' \
--header 'Content-Type: application/json' \
--data '{
    "latitude": 39.7456,
    "longitude": -97.0892
}'
```

And you should receive a response that looks like:

```json
{
    "shortForecast": "Chance Showers And Thunderstorms",
    "temperatureDescription": "moderate"
}
```