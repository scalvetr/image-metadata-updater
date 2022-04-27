package metadata

import (
	"context"
	"fmt"
	gphotos "github.com/gphotosuploader/google-photos-api-client-go/v2"
	"golang.org/x/oauth2"
	a "image-metadata-updater/album"
	c "image-metadata-updater/config"
	"log"
)

type Uploader struct {
	client *gphotos.Client
	config c.Config
}

func (u Uploader) init(config c.Config) {
	ctx := context.Background()
	oc := oauth2.Config{
		ClientID:     config.GoogleApi.ClientID,
		ClientSecret: config.GoogleApi.ClientSecret,
	}
	var token *oauth2.Token
	tc := oc.Client(ctx, token)
	var err error
	u.client, err = gphotos.NewClient(tc)
	if err != nil {
		log.Fatal(err)
	}
	u.config = config
}

func CreateAlbum(album a.AlbumInfo) {
	fmt.Println("Album: ", album.Year, album.Month, album.Name)
}

func UploadFile(album a.AlbumInfo, filepath string) {
	fmt.Println("file: ", filepath)
}
