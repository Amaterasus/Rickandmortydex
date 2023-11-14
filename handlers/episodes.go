package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/amaterasus/Rickandmortydex/api"
)

func RenderFullEpisodePageHTML(c *gin.Context) {

	name, nameExists := c.GetQuery("name")

	episodes, err := api.GetEpisodes("1", name)


	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}

	nextPage := 2

	if c.GetHeader("HX-Request") == "true"  && nameExists {
		if name == "" {
			c.Header("HX-Push-Url", "/episodes")
		}

		// due to my poor naming this looks identical to our other html document but it's just this is a pluralisation
		c.HTML(http.StatusOK, "episodes.html", gin.H{"episodes": episodes, "nextPage": nextPage, "length": len(episodes), "name": name})
		return
	}

	c.HTML(http.StatusOK, "episode.html", gin.H{"episodes": episodes, "nextPage": nextPage, "length": len(episodes), "name": name})
}

func RenderEpisodesHTML(c *gin.Context) {

	page := "1"
	s, ok := c.GetQuery("page")
	if ok {
		page = s
	}

	name, _ := c.GetQuery("name")

	episodes, err := api.GetEpisodes(page, name)

	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", nil)
	}

	nextPage, err := strconv.ParseInt(page, 10, 32)

	if err != nil {
		log.Println("handle this later!")
	}

	nextPage = nextPage + 1

	c.HTML(http.StatusOK, "episodes.html", gin.H{"episodes": episodes, "nextPage": nextPage, "length": len(episodes), "name": name})
}
