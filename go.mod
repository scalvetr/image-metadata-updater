module image-metadata-updater

go 1.18

require (
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	image-metadata-updater/action v0.0.0
	image-metadata-updater/album v0.0.0
	image-metadata-updater/config v0.0.0
	image-metadata-updater/metadata v0.0.0
)

require (
	github.com/dsoprea/go-exif/v3 v3.0.0-20210625224831-a6301f85c82b // indirect
	github.com/dsoprea/go-logging v0.0.0-20200517223158-a10564966e9d // indirect
	github.com/dsoprea/go-utility/v2 v2.0.0-20200717064901-2fccff4aa15e // indirect
	github.com/go-errors/errors v1.1.1 // indirect
	github.com/golang/geo v0.0.0-20200319012246-673a6f80352d // indirect
	golang.org/x/net v0.0.0-20200513185701-a91f0712d120 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	image-metadata-updater/action v0.0.0 => ./action
	image-metadata-updater/album v0.0.0 => ./album
	image-metadata-updater/config v0.0.0 => ./config
	image-metadata-updater/metadata v0.0.0 => ./metadata
)
