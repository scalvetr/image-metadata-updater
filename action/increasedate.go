package action

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	c "photo-manager-cli/config"
	"photo-manager-cli/metadata"
	"time"
)

func IncreaseDate(config c.Config) {
	fmt.Println("[INIT] IncreaseDate")
	fmt.Println("path: ", config.Path)
	fmt.Println("date_range_from: ", config.IncreaseDateConfig.DateRangeFrom)
	fmt.Println("date_range_to: ", config.IncreaseDateConfig.DateRangeTo)
	fmt.Println("increase_milliseconds: ", config.IncreaseDateConfig.IncreaseSeconds)

	files, err := os.ReadDir(config.Path)
	if err != nil {
		log.Fatal(err)
	}

	var directories []os.DirEntry
	for _, file := range files {
		if file.IsDir() {
			directories = append(directories, file)
		}
	}
	for _, directory := range directories {
		fmt.Println("[IncreaseDate] - ", directory)
		increaseSeconds(config.Path, directory.Name(), config.IncreaseDateConfig)
	}
	fmt.Println("[Finish] IncreaseDate")
}

func increaseSeconds(basePath string, directory string, config c.IncreaseDateConfig) {
	fmt.Println(fmt.Sprintf("- directory: %s", directory))
	filepath.Walk(filepath.Join(basePath, directory),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				fmt.Println(fmt.Sprintf(" - file %s mod_time: %s", info.Name(), info.ModTime()))
				if info.ModTime().After(config.DateRangeFrom) && info.ModTime().Before(config.DateRangeTo) {
					newTime := info.ModTime().Add(time.Second * time.Duration(config.IncreaseSeconds))
					fmt.Printf("   increase to %s\n", newTime)
					metadata.UpdateMetadataDate(path, info, &newTime, true, false, nil, nil)
					if err := os.Chtimes(path, newTime, newTime); err != nil {
						log.Fatal(err)
					}
				}
			}
			return nil
		})
}
