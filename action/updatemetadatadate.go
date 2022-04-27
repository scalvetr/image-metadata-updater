package action

import (
	"fmt"
	"image-metadata-updater/config"
	"image-metadata-updater/metadata"
	"log"
	"os"
	"path/filepath"
	"time"
)

func UpdateMetadataDate(config config.Config) {
	var fileDateTime *time.Time
	layout := "2006-01-02T15:04:05Z07:00"
	valueStr := config.Date
	k, err := time.Parse(layout, valueStr)

	if err != nil {
		log.Fatal(err)
	}
	fileDateTime = &k

	fmt.Println("DateTime", fileDateTime)
	processFixedDate(config.Path, fileDateTime)
}

func processFixedDate(basePath string, fileDateTime *time.Time) {
	filepath.Walk(basePath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			metadata.UpdateMetadataDate(path, info, fileDateTime)
			return nil
		})
}
