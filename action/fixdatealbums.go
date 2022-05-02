package action

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	a "photo-manager-cli/album"
	c "photo-manager-cli/config"
	"photo-manager-cli/metadata"
)

func FixDateAlbums(config c.Config) {
	fmt.Println("[INIT] FixDateAlbums")
	fmt.Println("path: ", config.Path)
	fmt.Println("albumNamePattern: ", config.AlbumInfoConfig.GetAlbumNamePattern())
	fmt.Println("getFolderRegexp: ", config.AlbumInfoConfig.GetFolderRegexp())

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
		var info = a.ExtractAlbumInfo(directory, config.AlbumInfoConfig)
		fmt.Println("[FixDateAlbums] - ", info.Year, info.Month, info.Name)
		fixDate(config.Path, directory, info)
	}
	fmt.Println("[Finish] FixDateAlbums")
}

func fixDate(basePath string, directory fs.FileInfo, albumInfo a.AlbumInfo) {
	filepath.Walk(filepath.Join(basePath, directory.Name()),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				metadata.FixDate(path, albumInfo.Year, albumInfo.Month)
			}
			return nil
		})
}
