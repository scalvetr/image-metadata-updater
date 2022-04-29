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
	fmt.Println("[INIT] UpdateMetadataDate")
	fmt.Println("path: ", config.Path)
	fmt.Println("date: ", config.UpdateMetadataDateConfig.Date)
	fmt.Println("override: ", config.UpdateMetadataDateConfig.Override)

	var fileDateTime *time.Time
	layout := "2006-01-02T15:04:05Z07:00"
	valueStr := config.UpdateMetadataDateConfig.Date
	k, err := time.Parse(layout, valueStr)
	if err != nil {
		log.Fatal(err)
	}
	fileDateTime = &k
	processUpdateMetadataDate(config.Path, fileDateTime, config.UpdateMetadataDateConfig.Override)
	fmt.Println("[Finish] UpdateMetadataDate")
}

func processUpdateMetadataDate(basePath string, fileDateTime *time.Time, override bool) {
	filepath.Walk(basePath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			metadata.UpdateMetadataDate(path, info, fileDateTime, override)
			return nil
		})
}
