package metadata

import (
	"fmt"
	config "image-metadata-updater/album"
)

func CreateAlbum(album config.AlbumInfo) {
	fmt.Println("Album: ", album.Year, album.Month, album.Name)
}

func UploadFile(album config.AlbumInfo, filepath string) {
	fmt.Println("file: ", filepath)
}
