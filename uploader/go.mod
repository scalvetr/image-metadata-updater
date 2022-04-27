module image-metadata-updater/uploader

go 1.18

require (
	github.com/gphotosuploader/google-photos-api-client-go/v2 v2.3.0
	golang.org/x/oauth2 v0.0.0-20220411215720-9780585627b5
	image-metadata-updater/album v0.0.0
	image-metadata-updater/config v0.0.0
)

require (
	github.com/gadelkareem/cachita v0.2.1 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/gphotosuploader/googlemirror v0.5.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.6.8 // indirect
	github.com/mediocregopher/radix/v3 v3.2.0 // indirect
	github.com/vmihailenco/msgpack v4.0.1+incompatible // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	google.golang.org/api v0.30.0 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)

replace (
	image-metadata-updater/album v0.0.0 => ../album
	image-metadata-updater/config v0.0.0 => ../config
)
