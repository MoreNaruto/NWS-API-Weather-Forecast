package models

type LatLongRequest struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}

type GridLocationResponse struct {
	Properties GridLocationProperties `json:"properties"`
}

type GridLocationProperties struct {
	CWA string `json:"cwa"`
	GridId string `json:"gridId"`
	ForecastUrl string `json:"forecast"`
}

type GridPointForecastResponse struct {
	Properties GridPointForecastProperties `json:"properties"`
}

type GridPointForecastProperties struct {
	Periods []GridPointForecastPeriod `json:"periods"`
}

type GridPointForecastPeriod struct {
	Temperature int `json:"temperature"`
	TemperatureUnit string `json:"temperatureUnit"`
	ShortForecast string `json:"shortForecast"`
	DetailedForecast string `json:"detailedForecast"`
}

type TodayForecastResponse struct {
	ShortForecast string `json:"shortForecast"`
	TemperatureDescription string `json:"temperatureDescription"`
}