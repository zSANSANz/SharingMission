package controllers

import (
	"net/http"
	"rumahbelajar-api/config"
	"rumahbelajar-api/models"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	_ "github.com/heroku/x/hmetrics/onload"
)

func PostWalikelas(c *gin.Context) {
	db := config.InitDb()
	defer db.Close()

	var walikelas models.WaliKelass
	c.Bind(&walikelas)

	if walikelas.Kelas != 0 && walikelas.Nama != "" {
		// INSERT INTO "walikelass" (name) VALUES (walikelas.Name);
		db.Create(&walikelas)
		// Display error
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"code":    200,
			"message": "Success",
			"data":    walikelas,
		})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/walikelass
}

func GetWalikelass(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	var walikelass []models.WaliKelass
	// SELECT * FROM walikelass
	db.Find(&walikelass)

	// Display JSON result
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"code":    200,
		"message": "Success",
		"data":    walikelass,
	})
	// curl -i http://localhost:8080/api/v1/walikelass
}

func GetWalikelas(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	id := c.Params.ByName("id")
	var walikelas models.WaliKelass
	// SELECT * FROM walikelass WHERE id = 1;
	db.First(&walikelas, id)

	if walikelas.ID != 0 {
		// Display JSON result
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":    true,
			"code":      200,
			"message":   "Success",
			"walikelas": walikelas,
		})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Walikelas not found"})
	}

	// curl -i http://localhost:8080/api/v1/walikelass/1
}

func UpdateWalikelas(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	// Get id walikelas
	id := c.Params.ByName("id")
	var walikelas models.WaliKelass
	// SELECT * FROM walikelass WHERE id = 1;
	db.First(&walikelas, id)

	if walikelas.Kelas != 0 && walikelas.Nama != "" {

		if walikelas.ID != 0 {
			var newWalikelas models.WaliKelass
			c.Bind(&newWalikelas)

			result := models.WaliKelass{
				Kelas: newWalikelas.Kelas,
				Nama:  newWalikelas.Nama,
			}

			// UPDATE walikelass SET firstname='newWalikelas.Firstname', lastname='newWalikelas.Lastname' WHERE id = walikelas.Id;
			db.Save(&result)
			// Display modified data in JSON message "success"
			c.JSON(200, gin.H{"success": result})
		} else {
			// Display JSON error
			c.JSON(404, gin.H{"error": "Walikelas not found"})
		}

	} else {
		// Display JSON error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/walikelass/1
}

func DeleteWalikelas(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	// Get id walikelas
	id := c.Params.ByName("id")
	var walikelas models.WaliKelass
	// SELECT * FROM walikelass WHERE id = 1;
	db.First(&walikelas, id)

	if walikelas.ID != 0 {
		// DELETE FROM walikelass WHERE id = walikelas.Id
		db.Delete(&walikelas)
		// Display JSON result
		c.JSON(200, gin.H{"success": "Walikelas #" + id + " deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Walikelas not found"})
	}

	// curl -i -X DELETE http://localhost:8080/api/v1/walikelass/1
}
