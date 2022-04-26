package action

import (
	"fmt"
	album "image-metadata-updater/album"
	"image-metadata-updater/config"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func UpdateDate(config config.Config) {

	files, err := ioutil.ReadDir(config.Path)
	if err != nil {
		log.Fatal(err)
	}

	var directories []fs.FileInfo
	for _, file := range files {
		if file.IsDir() {
			directories = append(directories, file)
		}
	}

	var fileDateTime *time.Time
	layout := "2006-01-02T15:04:05Z-07"
	valueStr := config.Date
	k, err := time.Parse(layout, valueStr+"Z+02")

	if err != nil {
		log.Fatal(err)
	}
	fileDateTime = &k

	fmt.Println("DateTime", fileDateTime)

	for _, directory := range directories {
		var info = album.ExtractAlbumInfo(directory)
		fmt.Println(info.Year, info.Month, info.Name)
		processFixedDate(config.Path, directory, fileDateTime)
	}
}

func processFixedDate(basePath string, directory fs.FileInfo, fileDateTime *time.Time) {
	filepath.Walk(filepath.Join(basePath, directory.Name()),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			err = os.Chtimes(path, *fileDateTime, *fileDateTime)
			if err != nil {
				fmt.Println(err)
			}
			return nil
		})
}
