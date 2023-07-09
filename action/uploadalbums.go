package action

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	a "photo-manager-cli/album"
	c "photo-manager-cli/config"
	u "photo-manager-cli/uploader"
)

func UploadAlbums(config c.Config) {
	fmt.Println("[INIT] UploadAlbums")
	fmt.Println("path: ", config.Path)
	fmt.Println("albumNamePattern: ", config.AlbumInfoConfig.GetAlbumNamePattern())
	fmt.Println("getFolderRegexp: ", config.AlbumInfoConfig.GetFolderRegexp())

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

	u, err := u.CreateUploader(config)
	if err != nil {
		log.Fatal(err)
	}

	for _, directory := range directories {
		var info = a.ExtractAlbumInfo(directory, config.AlbumInfoConfig)
		fmt.Println(info.Year, info.Month, info.Name)
		uploadAlbum(config.Path, directory, info, *u)
	}
	fmt.Println("[Finish] UploadAlbums")
}

func uploadAlbum(basePath string, directory os.DirEntry, albumInfo a.AlbumInfo, uploader u.Uploader) {
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
