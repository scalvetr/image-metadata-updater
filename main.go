package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	a "photo-manager-cli/action"
	c "photo-manager-cli/config"
)

func readConfig() []c.Config {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var data []c.Config
	err2 := yaml.Unmarshal(file, &data)
	if err2 != nil {
		log.Fatal(err2)
	}
	return data
}

func main() {
	configs := readConfig()

	for _, config := range configs {
		fmt.Println("------------------------------")
		if config.Action == c.UpdateDateFromMetadata {
			a.UpdateDateFromMetadata(config)
		} else if config.Action == c.UpdateMetadata {
			a.UpdateMetadataDate(config)
		} else if config.Action == c.UploadAlbums {
			a.UploadAlbums(config)
		} else if config.Action == c.CheckAlbumDateMismatch {
			a.CheckAlbumDateMismatch(config)
		} else if config.Action == c.IncreaseDate {
			a.IncreaseDate(config)
		}
	}
	fmt.Println("------------------------------")

}
