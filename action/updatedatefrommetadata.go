package action

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"photo-manager-cli/config"
	"photo-manager-cli/metadata"
)

func UpdateDateFromMetadata(config config.Config) {
	fmt.Println("[INIT] UpdateDateFromMetadata")
	fmt.Println("path: ", config.Path)

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

	for _, directory := range directories {
		fmt.Println(directory.Name())
		processMetadata(config.Path, directory)
	}
	fmt.Println("[Finish] UpdateDateFromMetadata")
}

func processMetadata(basePath string, directory fs.FileInfo) {
	filepath.Walk(filepath.Join(basePath, directory.Name()),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			metadata.UpdateDateFromMetadata(path, info)
			return nil
		})
}
