package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/amaterasus/Rickandmortydex/api"
)

func RenderFullEpisodePageHTML(c *gin.Context) {

	episodes, err := api.GetEpisodes("1")

	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", nil)
	}

	nextPage := 2

	c.HTML(http.StatusOK, "episode.html", gin.H{"episodes": episodes, "nextPage": nextPage, "length": len(episodes)})
}

func RenderEpisodesHTML(c *gin.Context) {

	page := "1"
	s, ok := c.GetQuery("page")
	if ok {
		page = s
	}

	episodes, err := api.GetEpisodes(page)

	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", nil)
	}

	nextPage, err := strconv.ParseInt(page, 10, 32)

	if err != nil {
		log.Println("handle this later!")
	}

	nextPage = nextPage + 1

	c.HTML(http.StatusOK, "episodes.html", gin.H{"episodes": episodes, "nextPage": nextPage, "length": len(episodes)})
}
