package main

import (
	"my_Music_App/Get_SongByGenre/routes"
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
	routes.SongGetRoute(r)
	// Start the server and listen on port 8081
	r.Run(":8081")
}
