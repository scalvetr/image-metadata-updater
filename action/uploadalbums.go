package action

import (
	"fmt"
	a "image-metadata-updater/album"
	c "image-metadata-updater/config"
	u "image-metadata-updater/uploader"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func UploadAlbums(config c.Config) {
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

	u, err := u.CreateUploader(config)
	if err != nil {
		log.Fatal(err)
	}

	for _, directory := range directories {
		var info = a.ExtractAlbumInfo(directory)
		fmt.Println(info.Year, info.Month, info.Name)
		uploadAlbum(config.Path, directory, info, *u)
	}
	fmt.Println("[Finish] UploadAlbums")
}

func uploadAlbum(basePath string, directory fs.FileInfo, albumInfo a.AlbumInfo, uploader u.Uploader) {
	albumId, err := uploader.CreateAlbum(albumInfo)
	if err != nil {
		log.Fatal(err)
	}
	filepath.Walk(filepath.Join(basePath, directory.Name()),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				uploader.UploadFile(path, albumId)
			}
			return nil
		})
}
