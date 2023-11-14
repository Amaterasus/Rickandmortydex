package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/amaterasus/Rickandmortydex/api"
)

func RenderFullLocationPageHTML(c *gin.Context) {

	name, nameExists := c.GetQuery("name")

	locations, err := api.GetLocations("1", name)


	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}

	nextPage := 2

	if c.GetHeader("HX-Request") == "true"  && nameExists {
		log.Println("HTMX request")

		if name == "" {
			c.Header("HX-Push-Url", "/locations")
		}

		// due to my poor naming this looks identical to our other html document but it's just this is a pluralisation
		c.HTML(http.StatusOK, "locations.html", gin.H{"locations": locations, "nextPage": nextPage, "length": len(locations), "name": name})
		return
	}

	c.HTML(http.StatusOK, "location.html", gin.H{"locations": locations, "nextPage": nextPage, "length": len(locations), "name": name})
}

func RenderLocationsHTML(c *gin.Context) {

	page := "1"
	s, ok := c.GetQuery("page")
	if ok {
		page = s
	}

	name, _ := c.GetQuery("name")

	locations, err := api.GetLocations(page, name)

	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", nil)
	}

	nextPage, err := strconv.ParseInt(page, 10, 32)

	if err != nil {
		log.Println("handle this later!")
	}

	nextPage = nextPage + 1

	c.HTML(http.StatusOK, "locations.html", gin.H{"locations": locations, "nextPage": nextPage, "length": len(locations), "name": name})
}
