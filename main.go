package main

import(
	"fmt"
	nasaApi "MarsWeatherApp/nasa_api"
	database "MarsWeatherApp/database"
)

func main(){
	client := database.Connection()
	test := nasaApi.GetData()

	client.SendData(test)
	fmt.Print("Success")
}