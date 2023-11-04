package controllers

import (
	"database/sql"
	"log"
	"my_Music_App/Get_SongByGenre/models"
	"my_Music_App/db_conn"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func GetSongByGenre(c *gin.Context) {
	songGenre := c.Param("genre")

	// Get the page number from the URL query parameter (default to page 1)
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	// Calculate the number of rows to skip based on the page number and the desired limit
	limit := 10
	offset := (page - 1) * limit

	// Initialize the database connection
	db, err := db_conn.InitDB()
	if err != nil {
		log.Printf("Error in Initializing Database :%s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error ": "Internal Server Error"})
		return
	}

	defer db.Close()

	// Check if the song with the specified genre exists in the database.
	var existingSong string
	err = db.QueryRow("SELECT genre FROM songInfo WHERE genre=?", songGenre).Scan(&existingSong)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Songs not found for the given genre"})
			return
		}
		log.Printf("Error in checking post existence: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Define the SQL query for retrieving all songs by genre
	query := "SELECT * FROM songInfo WHERE genre=?"
	rows, err := db.Query(query, songGenre)
	if err != nil {
		log.Printf("Error in retrieving songs by genre :%s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error ": err.Error()})
		return
	}

	defer rows.Close()

	var songs []models.SongInfo
	// Iterate through the rows and retrieve song information
	for rows.Next() {
		var song models.SongInfo
		// Scan the data from the current row into the 'song' variable
		err = rows.Scan(&song.ID, &song.Name, &song.Artists, &song.Genre, &song.PublishYear, &song.Language)
		if err != nil {
			log.Printf("Error in retrieving songs by genre :%s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error ": err.Error()})
			return
		}
		// Append the retrieved song to the 'songs' slice
		songs = append(songs, song)
	}

	if err := rows.Err(); err != nil {
		// Check if there was an error while iterating through rows
		log.Printf("Error in retrieving rows :%s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error ": err.Error()})
		return
	}

	// Determine the start and end index for the current page
	startIndex := offset
	endIndex := offset + limit

	// Ensure that the indices are within the bounds of the songs slice
	if startIndex < 0 {
		startIndex = 0
	}
	if endIndex > len(songs) {
		endIndex = len(songs)
	}

	// Slice the songs for the current page
	pageSongs := songs[startIndex:endIndex]

	// If there were no errors, return the retrieved songs for the current page as a JSON response with a 200 status code
	c.JSON(http.StatusOK, pageSongs)
}
