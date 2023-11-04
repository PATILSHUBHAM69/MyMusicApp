package main

import (
	"my_Music_App/Insert_Song/routes"
	"my_Music_App/db_conn"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Create a new Gin router with default
	r := gin.Default()
	// Initialize the database connection
	db_conn.InitDB()
	// Register the route for getting songs by genre
	routes.SongInsertRoute(r)
	// Start the server and listen on port 8080
	r.Run(":8080")
}
