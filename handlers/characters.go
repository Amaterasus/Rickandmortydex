package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/amaterasus/Rickandmortydex/api"
)

func RenderFullCharacterPageHTML(c *gin.Context) {

	characters, err := api.GetCharacters("1")

	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", nil)
	}

	nextPage := 2

	c.HTML(http.StatusOK, "character.html", gin.H{"characters": characters, "nextPage": nextPage, "length": len(characters)})
}


func RenderCharactersHTML(c *gin.Context) {

	page := "1"
	s, ok := c.GetQuery("page")
	if ok {
		page = s
	}

	characters, err := api.GetCharacters(page)

	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", nil)
	}

	nextPage, err := strconv.ParseInt(page, 10, 32)

	if err != nil {
		log.Println("handle this later!")
	}

	nextPage = nextPage + 1

	c.HTML(http.StatusOK, "characters.html", gin.H{"characters": characters, "nextPage": nextPage, "length": len(characters)})
}

func RenderMainCharactersBioHTML(c *gin.Context) {

	id := c.Param("id")

	character, episodes, err := api.GetCharacter(id)

	log.Println(character.Location)

	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", nil)
	}

	c.HTML(http.StatusOK, "fullCharacterBio.html", gin.H{"character": character, "episodes": episodes})
}
