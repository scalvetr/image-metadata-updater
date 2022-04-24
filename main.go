package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	album "image-metadata-updater/album"
	config "image-metadata-updater/config"
	metadata "image-metadata-updater/metadata"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func readConfig() config.Config {
	yfile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var data config.Config
	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {
		log.Fatal(err2)
	}
	return data
}

func main() {
	config := readConfig()
	files, err := ioutil.ReadDir(config.BasePath)
	if err != nil {
		log.Fatal(err)
	}

	var directories []fs.FileInfo
	for _, file := range files {
		if accept(file) {
			directories = append(directories, file)
		}
	}

	for _, directory := range directories {
		var info = album.ExtractAlbumInfo(directory)
		fmt.Println(info.Year, info.Month, info.Name)
		process(config.BasePath, directory, info)
	}
}

func process(basePath string, directory fs.FileInfo, albumInfo album.AlbumInfo) {
	filepath.Walk(filepath.Join(basePath, directory.Name()),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			metadata.ProcessMetadata(albumInfo, path, info)
			return nil
		})
}

func accept(file fs.FileInfo) bool {
	return file.IsDir()
}
