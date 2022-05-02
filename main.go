package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	a "photo-manager-cli/action"
	c "photo-manager-cli/config"
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
		} else if config.Action == c.UpdateMetadata {
			a.UpdateMetadataDate(config)
		} else if config.Action == c.UploadAlbums {
			a.UploadAlbums(config)
		} else if config.Action == c.FixDateAlbums {
			a.FixDateAlbums(config)
		}
	}
	fmt.Println("------------------------------")

}
