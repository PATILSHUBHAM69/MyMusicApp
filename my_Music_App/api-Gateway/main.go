package main

import (
	getroutes "my_Music_App/Get_SongByGenre/routes" // Import the routes from Get_SongByGenre
	insertroutes "my_Music_App/Insert_Song/routes"  // Import the routes from Insert_Song

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()

	// Initialize routes for each microservice
	getroutes.SongGetRoute(r)       // Add routes from Get_SongByGenre
	insertroutes.SongInsertRoute(r) // Add routes from Insert_Song

	// Start the Gin server on port 8082
	r.Run(":8082")
}
