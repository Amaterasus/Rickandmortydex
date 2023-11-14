package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/amaterasus/Rickandmortydex/api"
)

func RenderFullCharacterPageHTML(c *gin.Context) {

	name, nameExists := c.GetQuery("name")

	characters, err := api.GetCharacters("1", name)


	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}

	if c.GetHeader("HX-Request") == "true"  && nameExists {
		if name == "" {
			c.Header("HX-Push-Url", "/characters")
		}

		// due to my poor naming this looks identical to our other html document but it's just this is a pluralisation
		c.HTML(http.StatusOK, "characters.html", gin.H{"characters": characters, "nextPage": 2, "length": len(characters), "name": name})
		return
	}

	c.HTML(http.StatusOK, "character.html", gin.H{"characters": characters, "nextPage": 2, "length": len(characters), "name": name})
}


func RenderCharactersHTML(c *gin.Context) {

	page := "1"
	s, ok := c.GetQuery("page")
	if ok {
		page = s
	}
	name, _ := c.GetQuery("name")

	characters, err := api.GetCharacters(page, name)

	if err != nil {
		log.Println(err)
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}

	nextPage, err := strconv.ParseInt(page, 10, 32)

	if err != nil {
		log.Println("handle this later!")
	}

	nextPage = nextPage + 1

	c.HTML(http.StatusOK, "characters.html", gin.H{"characters": characters, "nextPage": nextPage, "length": len(characters), "name": name})
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
