package api

import (
	"strings"
)

func sliceIds(urls []string) (idString string) {

	var urlIds []string

	for _, url := range urls{
		splitUrl := strings.Split(url, "/")
		id := splitUrl[len(splitUrl)-1]

		urlIds = append(urlIds, id)
	}

	idString = strings.Join(urlIds, ",")

	return
}
