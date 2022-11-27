package api

import (
	"MarsWeatherApp/database"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func StartAPI() {
	app := fiber.New()                   // Creates fiber instance for api
	app.Get("/currentdata", currentData) // Endpoint and function to call
	app.Listen(":8084")                  // port number for api
}

func currentData(c *fiber.Ctx) error {
	client := database.Connection() // Fetch database connection/client
	msg := client.RetreiveData()    // Retrieve data from database using mongoDB drivers
	//result := make(map[string][]nasa_api.Soles)

	newData, _ := (json.Marshal(msg))    // method for converting a go struct into json
	return c.SendString(string(newData)) // Sends json object
}
