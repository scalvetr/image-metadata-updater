package metadata

import (
	"context"
	"fmt"
	gphotos "github.com/gphotosuploader/google-photos-api-client-go/v2"
	"github.com/gphotosuploader/google-photos-api-client-go/v2/albums"
	"github.com/gphotosuploader/google-photos-api-client-go/v2/media_items"
	"golang.org/x/oauth2"
	a "image-metadata-updater/album"
	c "image-metadata-updater/config"
)

type Uploader struct {
	client *gphotos.Client
	config c.Config
	ctx    context.Context
	albums map[string]*albums.Album
}

func CreateUploader(config c.Config) (*Uploader, error) {
	var uploader = Uploader{}
	error := uploader.init(config)
	if error != nil {
		return nil, error
	}
	return &uploader, nil
}
func (u Uploader) init(config c.Config) error {
	u.ctx = context.Background()
	oc := oauth2.Config{
		ClientID:     config.GoogleApi.ClientID,
		ClientSecret: config.GoogleApi.ClientSecret,
	}
	var token *oauth2.Token
	tc := oc.Client(u.ctx, token)
	var err error
	u.client, err = gphotos.NewClient(tc)
	if err != nil {
		return err
	}
	u.config = config
	return nil
}
func (u Uploader) CreateAlbum(album a.AlbumInfo) (string, error) {
	fmt.Println("[GOOGLE_PHOTO] Album: ", album.Year, album.Month, album.Name)

	gAlbum, err := u.client.Albums.Create(u.ctx, album.GetName())
	if err != nil {
		return "", err
	}
	u.albums[gAlbum.ID] = gAlbum
	return gAlbum.ID, nil
}

func (u Uploader) UploadFile(filepath string, albumId string) (string, error) {
	fmt.Println("[GOOGLE_PHOTO] File: ", filepath)
	gToken, err := u.client.Uploader.UploadFile(u.ctx, filepath)
	if err != nil {
		return "", err
	}
	gMediaItem, err := u.client.MediaItems.CreateToAlbum(u.ctx, albumId, media_items.SimpleMediaItem{UploadToken: gToken})
	if err != nil {
		return "", err
	}
	return gMediaItem.ID, nil
}
