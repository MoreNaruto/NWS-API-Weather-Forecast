package main

import (
	"net/http"
	"os"
	"encoding/json"

	"fmt"

	"github.com/gin-gonic/gin"
	"nws-api/models"
)

func main() {
	router := gin.Default()

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	router.POST("/today-forecast", handleForecastedWeather)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	
	// Start server
	router.Run(":" + port)
}

func handleForecastedWeather(c *gin.Context) {
	var req models.LatLongRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Body"})
		return
	}

	if !isValidLatitude(req.Latitude) || !isValidLongitude(req.Longitude) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Latitude must be between -90 and 90; Longitude must be between -180 and 180"})
		return
	}

	url := fmt.Sprintf("https://api.weather.gov/points/%.4f,%.4f", req.Latitude, req.Longitude)

	resp, err := http.Get(url)

	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"error": "API call to NSW weather has failed"})
		return
	}
	defer resp.Body.Close()

	var gridLocationResponse models.GridLocationResponse
	if err := json.NewDecoder(resp.Body).Decode(&gridLocationResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse NSW weather API response"})
		return
	}

	var gridPointForecastResponse models.GridPointForecastResponse

	forecastUrl := gridLocationResponse.Properties.ForecastUrl

	resp, err = http.Get(forecastUrl)
	if err := json.NewDecoder(resp.Body).Decode(&gridPointForecastResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse NSW weather API response"})
		return
	}

	if len(gridPointForecastResponse.Properties.Periods) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No forecast exist for today"})
		return
	}
	currentWeather := gridPointForecastResponse.Properties.Periods[0]
	var temperatureDescription string
	if currentWeather.TemperatureUnit == "F" {
		temperatureDescription = temperatureDescriptionInFahrenheit(currentWeather.Temperature)
	} else {
		temperatureDescription = temperatureDescriptionInCelsius(currentWeather.Temperature)
	}

	c.JSON(http.StatusOK, models.TodayForecastResponse{
		ShortForecast: currentWeather.ShortForecast,
		TemperatureDescription: temperatureDescription,
	})
}

func isValidLatitude(lat float64) bool {
	return lat >= -90 && lat <= 90
}

func isValidLongitude(lon float64) bool {
	return lon >= -180 && lon <= 180
}

func temperatureDescriptionInFahrenheit(temperature int) string {
	switch {
	case temperature < 50:
		return "cold"
	case temperature >= 50 && temperature <= 77:
		return "moderate"
	default:
		return "hot"
	}
}

func temperatureDescriptionInCelsius(temperature int) string {
	switch {
	case temperature < 10:
		return "cold"
	case temperature >= 10 && temperature <= 25:
		return "moderate"
	default:
		return "hot"
	}
}