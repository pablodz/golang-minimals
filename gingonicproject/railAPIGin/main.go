package main

import (
	"database/sql"
	"log"
	"net/http"

	"example.com/mymodule/gingonic/railAPIGin/dbutils"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type StationResource struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
}

func GetStation(c *gin.Context) {
	var station StationResource
	id := c.Param("station_id")
	err := DB.QueryRow("SELECT ID,NAME,CAST(OPENING_TIME as CHAR), CAST (CLOSING_TIME as CHAR) FROM station WHERE id=?", id).Scan(&station.ID, &station.Name, &station.OpeningTime, &station.ClosingTime)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"result": station,
		})
	}
}

func CreateStation(c *gin.Context) {
	var station StationResource
	err := c.BindJSON(&station)
	if err == nil {
		statement, _ := DB.Prepare("INSERT INTO station (NAME, OPENING_TIME, CLOSING_TIME) VALUES (?,?,?);")
		result, _ := statement.Exec(station.Name, station.OpeningTime, station.ClosingTime)
		if err == nil {
			newID, _ := result.LastInsertId()
			station.ID = int(newID)
			c.JSON(http.StatusOK, gin.H{
				"result": station,
			})
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func RemoveStation(c *gin.Context) {
	id := c.Param("station-id")
	statement, _ := DB.Prepare("DELETE FROM station WHERE id=?")
	_, err := statement.Exec(id)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.String(http.StatusOK, "")
	}
}

func main() {
	var err error
	DB, err = sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}
	dbutils.Initilialize(DB)
	r := gin.Default()
	// Add routes to REST verb
	r.GET("/v1/stations/:station_id", GetStation)
	r.POST("/v1/stations", CreateStation)
	r.DELETE("/v1/stations/:station_id", RemoveStation)
	r.Run(":8000")
}

// curl -X POST\
// 	http://localhost:8000/v1/stations \
// 	-H 'cache-control: no-cache' \
// 	-H 'content-type: application/json' \
// 	-d '{"name":"Brooklyn","opening_time":"8:12:00","closing_time":"18:23:00"}'
