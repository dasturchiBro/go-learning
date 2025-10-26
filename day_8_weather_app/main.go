/*
	Weather App
	API Key: 021cc4d180793e2374ab5de3dd5ee818
	https://openweathermap.org/current
	https://api.openweathermap.org/data/2.5/weather?q={city name}&appid={API key}

	We need three things to get data from an API: io, log, net/http

	First, we need to perform a get request: http.Get(--URL--).
	After checking for errors, it is time to read the body data using io.ReadAll(resp.Body)
	Then, it is time to print out the result.
	We gotta make sure to close the request by defer resp.Body.Close()
	

	https://api.openweathermap.org/data/2.5/weather?q=Tashkent&appid=021cc4d180793e2374ab5de3dd5ee818
*/

package main 

import (
	"fmt"

	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"time"
)

type InnerDataCoord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type InnerDataWeatherInfo struct {
	Id int `json:"id"`
	Main string `json:"main"`
	Description string `json:"description"`
	Icon string `json:"icon"`
}


type InnerDataMain struct {
	Temp float64 `json:"temp"`
	Feels_like float64 `json:"feels_like"`
	Temp_min float64 `json:"temp_min"`
	Temp_max float64 `json:"temp_max"`
	Pressure int `json:"pressure"`
	Humidity int `json:"humidity"`
	Sea_level int `json:"sea_level"`
	Grnd_level int `json:"grnd_level"`
}

type InnerDataWind struct {
	Speed float64 `json:"speed"`
	Deg int `json:"deg"`
}

type InnerDataClouds struct {
	All float64 `json:"all"`
}

type InnerDataSys struct {
	Type int `json:"type"`
	Id int `json:"id"`
	Country string `json:"country"`
	Sunrise int `json:"sunrise"`
	Sunset int `json:"sunset"`
}


type WeatherData struct {
	Coord InnerDataCoord `json:"coord"`
	Weather []InnerDataWeatherInfo `json:"weather"`
	Base string `json:"base"`
	Main InnerDataMain `json:"main"`
	Visibility int `json:"visibility"`
	Wind InnerDataWind `json:"wind"`
	Clouds InnerDataClouds `json:"clouds"`
	Dt int `json:"dt"`
	Sys InnerDataSys `json:"sys"`
	Timezone int `json:"timezone"`
	Id int `json:"id"`
	Name string `json:"name"`
	Cod int `json:"cod"`
}

func get_info(URL string) (WeatherData, bool) {
	var weatherdata WeatherData
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return weatherdata, false
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	err = json.Unmarshal(body, &weatherdata)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return weatherdata, true
}

func KelvinToCelsius(value float64) float64 {
	return (value - 273.15)
}

func loader(t time.Duration) {
	fmt.Println("\n.")
	time.Sleep(t * time.Millisecond)
	fmt.Print("."); time.Sleep((t) * time.Millisecond); fmt.Print(".\n");time.Sleep((t) * time.Millisecond)
	fmt.Print("."); time.Sleep((t) * time.Millisecond);fmt.Print("."); time.Sleep((t) * time.Millisecond); fmt.Print(".\n");time.Sleep((t) * time.Millisecond)
}

func main() {
	for true {
		fmt.Print("***Welcome to real-time weather app.***\n\nEnter the name of the city: ")
		var city_name string  
		fmt.Scan(&city_name)
		loader(300)
		URL := "https://api.openweathermap.org/data/2.5/weather?q="+city_name+"&appid=021cc4d180793e2374ab5de3dd5ee818"
		weatherdata, isFound := get_info(URL)
		if isFound {
			temp_in_kelvin := weatherdata.Main.Temp
			feels_like := weatherdata.Main.Feels_like
			fmt.Printf("***Here is the information: \n\tCity: %s\n\tTemperature: %.f\n\tFeels like: %.f\n\tHumidity: %d\n\tSea Level: %d\n\tWeather: %v", city_name, KelvinToCelsius(temp_in_kelvin), KelvinToCelsius(feels_like), weatherdata.Main.Humidity, weatherdata.Main.Sea_level, weatherdata.Weather[0].Main)
			loader(300)
			fmt.Println("\n")
		} else {
			fmt.Println("Place not found. Please try again.\n\n")
		}
	}
}