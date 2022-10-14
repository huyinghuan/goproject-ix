package bundle

import (
	"embed"
	"log"
)

//go:embed asserts
var configFile embed.FS

func ReadFile(filename string) []byte {
	body, err := configFile.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
