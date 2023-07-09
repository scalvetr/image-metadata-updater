module photo-manager-cli/uploader

go 1.23

require (
	github.com/gphotosuploader/google-photos-api-client-go/v2 v2.4.1
	golang.org/x/oauth2 v0.10.0
	photo-manager-cli/album v0.0.0
	photo-manager-cli/config v0.0.0
)

require (
	cloud.google.com/go v0.110.2 // indirect
	cloud.google.com/go/compute v1.20.1 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/gadelkareem/cachita v0.2.3 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/gphotosuploader/googlemirror v0.5.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.4 // indirect
	github.com/mediocregopher/radix/v3 v3.8.1 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/api v0.130.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	photo-manager-cli/album v0.0.0 => ../album
	photo-manager-cli/config v0.0.0 => ../config
)
