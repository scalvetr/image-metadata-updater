package action

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"photo-manager-cli/config"
	"photo-manager-cli/metadata"
	"time"
)

func UpdateMetadataDate(config config.Config) {
	fmt.Println("[INIT] UpdateMetadata")
	fmt.Println("path: ", config.Path)
	fmt.Println("date: ", config.UpdateMetadataDateConfig.Date)
	fmt.Println("override: ", config.UpdateMetadataDateConfig.Override)

	processUpdateMetadataDate(config.Path, config.UpdateMetadataDateConfig)
	fmt.Println("[Finish] UpdateMetadata")
}

func processUpdateMetadataDate(basePath string, config config.UpdateMetadataConfig) {
	valueStr := config.Date
	override := config.Override
	var fileDateTime *time.Time
	if valueStr != "" {
		layout := "2006-01-02T15:04:05Z07:00"
		k, err := time.Parse(layout, valueStr)
		if err != nil {
			log.Fatal(err)
		}
		fileDateTime = &k
	}
	replaces := make(map[string]string)
	for _, replace := range config.Replace {
		replaces[replace.Day] = replace.NewDay
	}

	filepath.Walk(basePath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			metadata.UpdateMetadataDate(path, info, fileDateTime, override, replaces)
			return nil
		})
}
