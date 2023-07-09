module photo-manager-cli

go 1.23

require (
	gopkg.in/yaml.v3 v3.0.1
	photo-manager-cli/action v0.0.0
	photo-manager-cli/config v0.0.0
)

require (
	cloud.google.com/go/compute v1.20.1 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/dsoprea/go-exif/v3 v3.0.1 // indirect
	github.com/dsoprea/go-iptc v0.0.0-20200610044640-bc9ca208b413 // indirect
	github.com/dsoprea/go-jpeg-image-structure/v2 v2.0.0-20221012074422-4f3f7e934102 // indirect
	github.com/dsoprea/go-logging v0.0.0-20200710184922-b02d349568dd // indirect
	github.com/dsoprea/go-photoshop-info-format v0.0.0-20200610045659-121dd752914d // indirect
	github.com/dsoprea/go-utility/v2 v2.0.0-20221003172846-a3e1774ef349 // indirect
	github.com/gadelkareem/cachita v0.2.3 // indirect
	github.com/go-errors/errors v1.4.2 // indirect
	github.com/go-xmlfmt/xmlfmt v1.1.2 // indirect
	github.com/golang/geo v0.0.0-20230421003525-6adc56603217 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/gphotosuploader/google-photos-api-client-go/v2 v2.4.1 // indirect
	github.com/gphotosuploader/googlemirror v0.5.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.4 // indirect
	github.com/mediocregopher/radix/v3 v3.8.1 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/oauth2 v0.10.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/api v0.130.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	photo-manager-cli/album v0.0.0 // indirect
	photo-manager-cli/metadata v0.0.0 // indirect
	photo-manager-cli/uploader v0.0.0 // indirect
)

replace (
	photo-manager-cli/action v0.0.0 => ./action
	photo-manager-cli/album v0.0.0 => ./album
	photo-manager-cli/config v0.0.0 => ./config
	photo-manager-cli/metadata v0.0.0 => ./metadata
	photo-manager-cli/uploader v0.0.0 => ./uploader
)
