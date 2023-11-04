package routes

import (
	controllers "my_Music_App/Insert_Song/Controllers"

	"github.com/gin-gonic/gin"
)

// SongInsertRoute registers a route for insertinging songs.
func SongInsertRoute(r *gin.Engine) {
	r.POST("/insertNewSong", controllers.InsertNewSong)
}
