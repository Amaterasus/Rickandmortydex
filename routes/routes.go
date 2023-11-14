package routes

import (
	"github.com/amaterasus/Rickandmortydex/handlers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	r.SetTrustedProxies(nil)

	r.GET("/", handlers.RenderFullHomePageHTML)
	r.NoRoute(handlers.RenderNotFoundHTML)

	characters := r.Group("characters")
	characters.GET("", handlers.RenderFullCharacterPageHTML)
	characters.GET("/:id", handlers.RenderMainCharactersBioHTML)
	characters.POST("", handlers.RenderCharactersHTML)

	locations := r.Group("locations")
	locations.GET("", handlers.RenderFullLocationPageHTML)
	locations.POST("", handlers.RenderLocationsHTML)

	episodes := r.Group("episodes")
	episodes.GET("", handlers.RenderFullEpisodePageHTML)
	episodes.POST("", handlers.RenderEpisodesHTML)
}
