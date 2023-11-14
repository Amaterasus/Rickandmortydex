package routes

import (
	"github.com/amaterasus/Rickandmortydex/handlers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	r.SetTrustedProxies(nil)

	r.GET("/", handlers.RenderFullHomePageHTML)
	r.NoRoute(handlers.RenderNotFoundHTML)
}
