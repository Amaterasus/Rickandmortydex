package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderFullHomePageHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}
