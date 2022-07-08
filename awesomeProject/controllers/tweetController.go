package controllers

import (
	"awesomeProject/database"
	"awesomeProject/models"
	"encoding/json"
	"fmt"
	fiber2 "github.com/gofiber/fiber/v2"
	"io/ioutil"
)

func InitTweetsToDb(c *fiber2.Ctx) error {
	var tweet models.Tweet
	result := database.DB.First(&tweet)

	if result.RowsAffected > 0 {
		fmt.Println("Retreived")
		//return c.SendString("Retreived")
		c.JSON(models.InitRes{"Retreived"})
	}

	jsonFile, err := ioutil.ReadFile("elonmusk.json") // just pass the file name
	if err != nil {
		fmt.Print(err)
		fmt.Println("Failed")
		c.SendString("Failed")
	}

	data := []models.Tweet{}
	json.Unmarshal([]byte(string(jsonFile)), &data)

	for i, tweet := range data {
		database.DB.Create(&tweet)
		if i > 100 {
			break
		}
	}
	fmt.Println("Initialized")
	return c.SendString("Initialized")
}
