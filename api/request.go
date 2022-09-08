package api

import (
	"MarsWeatherApp/database"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func StartAPI() {
	client := database.Connection()								// Fetch database connection/client
	app := fiber.New()											// Creates fiber instance for api					
	// GET /api/register


	app.Get("/currentdata", func(c *fiber.Ctx) error {			// Endpoint for most recent information in database
		msg := client.RetreiveData()		 			 	    // Retrieve data from database using mongoDB drivers
		//result := make(map[string][]nasa_api.Soles)
		// result["data"] = msg
		newData, _ := (json.Marshal(msg))						// method for converting a go struct into json
		return c.SendString(string(newData)) 					// Sends json object 
	})



	// app.Post("/todos", func(c *fiber.Ctx) error {
	// 	date := c.Query("")
	// 	client.SendData(date)
	// 	return c.SendString("Request Successfully Posted")
	// })

	app.Listen(":9009")			// port number for api
}
