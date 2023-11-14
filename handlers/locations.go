package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/amaterasus/Rickandmortydex/api"
)

func RenderFullLocationPageHTML(c *gin.Context) {

	locations, err := api.GetLocations("1")

	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", nil)
	}

	nextPage := 2

	c.HTML(http.StatusOK, "location.html", gin.H{"locations": locations, "nextPage": nextPage, "length": len(locations)})
}

func RenderLocationsHTML(c *gin.Context) {

	page := "1"
	s, ok := c.GetQuery("page")
	if ok {
		page = s
	}

	locations, err := api.GetLocations(page)

	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", nil)
	}

	nextPage, err := strconv.ParseInt(page, 10, 32)

	if err != nil {
		log.Println("handle this later!")
	}

	nextPage = nextPage + 1

	c.HTML(http.StatusOK, "locations.html", gin.H{"locations": locations, "nextPage": nextPage, "length": len(locations)})
}
