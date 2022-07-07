package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Country struct {
	Name       string   `json:"name"`
	Code       string   `json:"code"`
	Capital    string   `json:"capital"`
	Area       int      `json:"area"`
	Population int      `json:"population"`
	Currency   string   `json:"currency"`
	Languages  []string `json:"languages"`
	Timezones  []string `json:"timezones"`
	Alpha2Code string   `json:"alpha2Code"`
	Alpha3Code string   `json:"alpha3Code"`
}

var countries = []Country{
	{
		Name:       "Afghanistan",
		Code:       "AF",
		Capital:    "Kabul",
		Area:       652230,
		Population: 27657145,
		Currency:   "AFN",
		Languages:  []string{"Pashto", "Dari"},
		Timezones:  []string{"UTC+04:30"},
		Alpha2Code: "AF",
		Alpha3Code: "AFG",
	},
	{
		Name:       "Albania",
		Code:       "AL",
		Capital:    "Tirana",
		Area:       287500,
		Population: 2886026,
		Currency:   "ALL",
		Languages:  []string{"Albanian"},
		Timezones:  []string{"UTC+01:00"},
		Alpha2Code: "AL",
		Alpha3Code: "ALB",
	},
	{
		Name:       "Algeria",
		Code:       "DZ",
		Capital:    "Algiers",
		Area:       2381740,
		Population: 37100000,
		Currency:   "DZD",
		Languages:  []string{"Arabic"},
		Timezones:  []string{"UTC+01:00"},
		Alpha2Code: "DZ",
		Alpha3Code: "DZA",
	},
	{
		Name:       "Andorra",
		Code:       "AD",
		Capital:    "Andorra la Vella",
		Area:       468,
		Population: 76965,
		Currency:   "EUR",
		Languages:  []string{"Catalan"},
		Timezones:  []string{"UTC+01:00"},
		Alpha2Code: "AD",
		Alpha3Code: "AND",
	},
	{
		Name:       "Angola",
		Code:       "AO",
		Capital:    "Luanda",
		Area:       1246700,
		Population: 19618432,
		Currency:   "AOA",
		Languages:  []string{"Portuguese"},
		Timezones:  []string{"UTC+01:00"},
		Alpha2Code: "AO",
		Alpha3Code: "AGO",
	},
	{
		Name:       "Antigua and Barbuda",
		Code:       "AG",
		Capital:    "Saint John's",
		Area:       442,
		Population: 86295,
		Currency:   "XCD",
		Languages:  []string{"English"},
		Timezones:  []string{"UTC-04:00"},
		Alpha2Code: "AG",
		Alpha3Code: "ATG",
	},
	{
		Name:       "Argentina",
		Code:       "AR",
		Capital:    "Buenos Aires",
		Area:       2766890,
		Population: 43593400,
		Currency:   "ARS",
		Languages:  []string{"Spanish"},
		Timezones:  []string{"UTC-03:00"},
		Alpha2Code: "AR",
		Alpha3Code: "ARG",
	},
	{
		Name:       "Armenia",
		Code:       "AM",
		Capital:    "Yerevan",
		Area:       29800,
		Population: 2968000,
		Currency:   "AMD",
		Languages:  []string{"Armenian"},
		Timezones:  []string{"UTC+04:00"},
		Alpha2Code: "AM",
		Alpha3Code: "ARM",
	},
	{
		Name:       "Australia",
		Code:       "AU",
		Capital:    "Canberra",
		Area:       7686850,
		Population: 24117344,
		Currency:   "AUD",
		Languages:  []string{"English", "Northern"},
		Timezones:  []string{"UTC+10:00"},
		Alpha2Code: "AU",
		Alpha3Code: "AUS",
	},
	{
		Name:       "Austria",
		Code:       "AT",
		Capital:    "Vienna",
		Area:       83870,
		Population: 8205000,
		Currency:   "EUR",
		Languages:  []string{"German"},
		Timezones:  []string{"UTC+01:00"},
		Alpha2Code: "AT",
		Alpha3Code: "AUT",
	},
	{
		Name:       "Azerbaijan",
		Code:       "AZ",
		Capital:    "Baku",
		Area:       86600,
		Population: 9459000,
		Currency:   "AZN",
		Languages:  []string{"Azerbaijani"},
		Timezones:  []string{"UTC+04:00"},
		Alpha2Code: "AZ",
		Alpha3Code: "AZE",
	},
}

func getCountries(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, countries)
}

func getCountry(ctx *gin.Context) {
	country := ctx.Param("country")
	for _, c := range countries {
		if c.Code == country {
			ctx.IndentedJSON(http.StatusOK, c)
			return
		}
	}
}

func addCountry(c *gin.Context) {
	var newCountry Country
	if err := c.BindJSON(&newCountry); err != nil {
		return
	}
	countries = append(countries, newCountry)
	c.IndentedJSON(http.StatusCreated, countries)
}

func updateCountry(ctx *gin.Context) {
	country := ctx.Param("country")
	for i, c := range countries {
		if c.Code == country {
			var newCountry Country
			if err := ctx.BindJSON(&newCountry); err != nil {
				return
			}
			countries[i] = newCountry
			ctx.IndentedJSON(http.StatusOK, countries)
			return
		}
	}
}

func removeCountry(ctx *gin.Context) {
	country := ctx.Param("country")
	for idx, c := range countries {
		if c.Code == country {
			countries = append(countries[:idx], countries[idx+1:]...)
			ctx.IndentedJSON(http.StatusOK, countries)
			return
		}
	}
}

func main() {

	router := gin.Default()

	router.GET("/countries", getCountries)
	router.GET("/countries/:country", getCountry)
	router.POST("/countries", addCountry)
	router.PUT("/countries/:country", updateCountry)
	router.PATCH("/countries/:country", updateCountry)
	router.DELETE("/countries/:country", removeCountry)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
