package controllers

import (
	"database/sql"
	"log"
	"my_Music_App/Insert_Song/models"
	"my_Music_App/db_conn"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func InsertNewSong(c *gin.Context) {
	// Initialize the database connection
	db, err := db_conn.InitDB()
	if err != nil {
		log.Printf("Error in Initializing Database :%s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error ": "Internal Server Error"})
		return
	}

	defer db.Close()

	//Bind the json data to newSong struct
	var newSong models.SongInfo
	if err := c.ShouldBindJSON(&newSong); err != nil {
		log.Printf("Error in binding song data :%s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}

	// Define SQL query for inserting new song info
	query := `INSERT INTO songInfo(name,artists,genre,publishyear,language) VALUES (?,?,?,?,?)`

	// Prepare and Execute the SQL statement
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Printf("Error in preparing SQL statement :%s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error ": err.Error()})
		return
	}

	defer stmt.Close()

	// Execute the SQL statement
	_, err = stmt.Exec(newSong.Name, newSong.Artists, newSong.Genre, newSong.PublishYear, newSong.Language)
	if err != nil {
		log.Printf("Error in Inserting new song info :%s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error ": "Internal Server Error"})
		return
	}

	// If there were no errors, return the Insert songs as a JSON response with a 200 status code
	c.JSON(http.StatusOK, gin.H{"message": "New Song Added Successfully"})

}
