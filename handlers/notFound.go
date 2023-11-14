package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderNotFoundHTML(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
