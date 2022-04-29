package metadata

import (
	"context"
	"fmt"
	gphotos "github.com/gphotosuploader/google-photos-api-client-go/v2"
	"github.com/gphotosuploader/google-photos-api-client-go/v2/albums"
	"github.com/gphotosuploader/google-photos-api-client-go/v2/media_items"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	a "photo-manager-cli/album"
	c "photo-manager-cli/config"
)

type Uploader struct {
	client *gphotos.Client
	config c.Config
	ctx    context.Context
	albums map[string]*albums.Album
}

func CreateUploader(config c.Config) (*Uploader, error) {
	var uploader *Uploader
	uploader = new(Uploader)
	error := uploader.init(config)
	if error != nil {
		return nil, error
	}
	return uploader, nil
}

// Requests a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func (u *Uploader) init(config c.Config) error {
	u.ctx = context.Background()

	b, err := ioutil.ReadFile("google_client.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	// If modifying these scopes, delete your previously saved token.json.
	oc, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/photoslibrary")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	token := getTokenFromWeb(oc)

	tc := oc.Client(u.ctx, token)
	client, err := gphotos.NewClient(tc)
	if err != nil {
		return err
	}
	u.client = client
	u.config = config
	u.albums = make(map[string]*albums.Album)
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
