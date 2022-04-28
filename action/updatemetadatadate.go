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
	fmt.Println("date: ", config.Date)

	var fileDateTime *time.Time
	layout := "2006-01-02T15:04:05Z07:00"
	valueStr := config.Date
	k, err := time.Parse(layout, valueStr)

	if err != nil {
		log.Fatal(err)
	}
	fileDateTime = &k
	processUpdateMetadataDate(config.Path, fileDateTime)
	fmt.Println("[Finish] UpdateMetadataDate")
}

func processUpdateMetadataDate(basePath string, fileDateTime *time.Time) {
	filepath.Walk(basePath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			metadata.UpdateMetadataDate(path, info, fileDateTime)
			return nil
		})
}
