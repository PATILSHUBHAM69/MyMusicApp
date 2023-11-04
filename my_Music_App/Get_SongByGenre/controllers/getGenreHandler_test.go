package controllers

import (
	"encoding/json"
	"fmt"
	"my_Music_App/Get_SongByGenre/models"
	"my_Music_App/db_conn"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func TestGetSongByGenre(t *testing.T) {
	// Create a temporary database for testing
	testDB, err := db_conn.InitDB()
	if err != nil {
		t.Fatalf("Error initializing test database: %v", err)
	}
	defer testDB.Close()

	// Initialize a Gin router for testing
	router := gin.Default()

	// Insert test data into the database
	testGenre := "Test Genre"
	_, err = testDB.Exec("INSERT INTO songInfo(name, artists, genre, publishyear, language) VALUES (?, ?, ?, ?, ?)",
		"Test Song 1", "Test Artist 1", testGenre, 2023, "English")
	if err != nil {
		t.Fatalf("Error inserting test data: %v", err)
	}

	// Create a request for the GetSongByGenre function
	req, err := http.NewRequest("GET", fmt.Sprintf("/songs/genre/%s", testGenre), nil)
	if err != nil {
		t.Fatalf("Error creating HTTP request: %v", err)
	}

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Set the database connection for the test
	db = testDB

	// Handle the request using the GetSongByGenre function
	router.GET("/songs/genre/:genre", GetSongByGenre)
	router.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, recorder.Code)
	}

	// Check the response body and unmarshal it into a slice of SongInfo
	var songs []models.SongInfo
	err = json.Unmarshal(recorder.Body.Bytes(), &songs)
	if err != nil {
		t.Fatalf("Error unmarshaling JSON response: %v", err)
	}

	// Check that at least one song is retrieved
	if len(songs) == 0 {
		t.Errorf("Expected at least one song, got none")
	}
}
