package weather_controller

import (
	"interface-testing/api/domain/weather_domain"
	"interface-testing/api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetWeather(c *gin.Context) {
	long, _ := strconv.ParseFloat(c.Param("longitude"), 64)
	lat, _ := strconv.ParseFloat(c.Param("latitude"), 64)
	request := weather_domain.WeatherRequest{
		ApiKey:    c.Param("apiKey"),
		Latitude:  lat,
		Longitude: long,
	}
	result, apiError := services.WeatherService.GetWeather(request)
	if apiError != nil {
		c.JSON(apiError.Status(), apiError)
		return
	}
	c.JSON(http.StatusOK, result)
}
