package database

import (
	"context"
	"fmt"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"MarsWeatherApp/nasa_api"
)


type MongoDb struct {
	Client *mongo.Client
}

//new connection
func Connection() *MongoDb {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://admin:****@domsdb.agpuaxn.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// test := MongoDb{Client: client, Context: &ctx}
	// fmt.Println(test.RetreiveData("shoes"))
	return &MongoDb{Client: client}
}

func (connection *MongoDb) SendData(info nasa_api.Soles) {
	
	db := connection.Client.Database("Mars_Application")
	coll := db.Collection("Data")

	docs := bson.M{"_id": info.Id, "Terrestrial_date": info.Terrestrial_date, "Sol": info.Sol, "Season": info.Season, 
					"Min_temp": info.Min_temp, "Max_temp": info.Max_temp,"Pressure": info.Pressure, 
					"Pressure_string": info.Pressure_string, "Atmo_opacity": info.Atmo_opacity, "Sunrise": info.Sunrise,
					"Sunset": info.Sunset, "Local_uv_irradiance_index": info.Local_uv_irradiance_index, "Min_gts_temp": info.Min_gts_temp, 
					"Max_gts_temp": info.Max_gts_temp}

	result, err := coll.InsertOne(context.TODO(), docs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("inserted document with ID %v\n", result.InsertedID)
}

// func (connection *MongoDb) RetreiveData(currentdate string) []nasa_api.Soles {
// 	var info nasa_api.Soles
// 	db := connection.Client.Database("Mars_Application") //Set Database
// 	coll := db.Collection("Data")                           //Set Collection
// 	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

// 	fmt.Println("Retreiving information...")
// 	filter := bson.M{"category": category} //Set Filter

// 	cursor, err := coll.Find(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Println("error reached")
// 		panic(err)
// 	}

// 	fmt.Println("Obtained data")

// 	for cursor.Next(context.TODO()) { //loop through all document
// 		var result nasa_api.Soles
// 		if err := cursor.Decode(&result); err != nil {
// 			panic(err)
// 		}
// 		info.Name = result.Name
// 		info.Number = result.Number
// 		infoArray = append(infoArray, info)
// 	}

// 	fmt.Println("Went through data")
// 	if err := cursor.Err(); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Successfully Retrieved")
// 	return []info

// }
