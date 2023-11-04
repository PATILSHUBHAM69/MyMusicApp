package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"my_Music_App/Insert_Song/models"
	"my_Music_App/db_conn"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func TestInsertNewSong(t *testing.T) {
	// Create a temporary database for testing
	testDB, err := db_conn.InitDB()
	if err != nil {
		t.Fatalf("Error initializing test database: %v", err)
	}
	defer testDB.Close()

	router := gin.Default()

	// Create a request payload
	newSong := models.SongInfo{
		Name:        "Test Song",
		Artists:     "Test Artist",
		Genre:       "Test Genre",
		PublishYear: 2023,
		Language:    "English",
	}
	payload, err := json.Marshal(newSong)
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}

	// Create an HTTP request with the JSON payload
	req, err := http.NewRequest("POST", "/insert", bytes.NewReader(payload))
	if err != nil {
		t.Fatalf("Error creating HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Set the database connection for the test
	db = testDB

	// Handle the request using the InsertNewSong function
	router.POST("/insert", InsertNewSong)
	router.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, recorder.Code)
	}

	// Check the response body
	expectedResponse := gin.H{"message": "New Song Added Successfully"}
	responseBody, _ := ioutil.ReadAll(recorder.Body)
	var response gin.H
	if err := json.Unmarshal(responseBody, &response); err != nil {
		t.Fatalf("Error unmarshaling JSON response: %v", err)
	}
	if !reflect.DeepEqual(expectedResponse, response) {
		t.Errorf("Expected response %v, got %v", expectedResponse, response)
	}
}
