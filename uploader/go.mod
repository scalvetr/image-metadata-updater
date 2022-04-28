module photo-manager-cli/uploader

go 1.18

require (
	github.com/gphotosuploader/google-photos-api-client-go/v2 v2.3.0
	golang.org/x/oauth2 v0.0.0-20220411215720-9780585627b5
	photo-manager-cli/album v0.0.0
	photo-manager-cli/config v0.0.0
)

require (
	cloud.google.com/go v0.100.2 // indirect
	cloud.google.com/go/compute v1.6.1 // indirect
	github.com/gadelkareem/cachita v0.2.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gphotosuploader/googlemirror v0.5.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.6.8 // indirect
	github.com/mediocregopher/radix/v3 v3.2.0 // indirect
	github.com/vmihailenco/msgpack v4.0.1+incompatible // indirect
	golang.org/x/net v0.0.0-20220425223048-2871e0cb64e4 // indirect
	google.golang.org/api v0.75.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace (
	photo-manager-cli/album v0.0.0 => ../album
	photo-manager-cli/config v0.0.0 => ../config
)
