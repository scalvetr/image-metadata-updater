package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	a "image-metadata-updater/action"
	c "image-metadata-updater/config"
	"io/ioutil"
	"log"
)

func readConfig() []c.Config {
	yfile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var data []c.Config
	err2 := yaml.Unmarshal(yfile, &data)
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
		} else if config.Action == c.UpdateMetadataDate {
			a.UpdateMetadataDate(config)
		} else if config.Action == c.UploadAlbums {
			a.UploadAlbums(config)
		}
	}
	fmt.Println("------------------------------")

}
