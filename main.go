package main

import (
	"MarsWeatherApp/nasa_api"
	"fmt"
)

func main() {
	data := nasa_api.GetData()
	fmt.Print(data)
}
