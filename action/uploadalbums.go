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
)

func UploadAlbums(config config.Config) {
	fmt.Println("[INIT] UploadAlbums")
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
		var info = album.ExtractAlbumInfo(directory)
		fmt.Println(info.Year, info.Month, info.Name)
		uploadAlbum(config.Path, directory, info)
	}
	fmt.Println("[Finish] UploadAlbums")
}

func uploadAlbum(basePath string, directory fs.FileInfo, albumInfo album.AlbumInfo) {
	filepath.Walk(filepath.Join(basePath, directory.Name()),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(albumInfo.Name, path)
			return nil
		})
}
