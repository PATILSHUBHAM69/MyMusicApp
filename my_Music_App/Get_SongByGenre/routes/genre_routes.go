package routes

import (
	controllers "my_Music_App/Get_SongByGenre/controllers"

	"github.com/gin-gonic/gin"
)

// SongGetRoute registers a route for retrieving songs by genre.
func SongGetRoute(r *gin.Engine) {
	r.GET("/getSongByGenre/:genre", controllers.GetSongByGenre)
}
