module photo-manager-cli

go 1.18

require (
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	photo-manager-cli/action v0.0.0
	photo-manager-cli/config v0.0.0
)

require (
	github.com/dsoprea/go-exif/v3 v3.0.0-20210625224831-a6301f85c82b // indirect
	github.com/dsoprea/go-iptc v0.0.0-20200609062250-162ae6b44feb // indirect
	github.com/dsoprea/go-jpeg-image-structure/v2 v2.0.0-20210512043942-b434301c6836 // indirect
	github.com/dsoprea/go-logging v0.0.0-20200710184922-b02d349568dd // indirect
	github.com/dsoprea/go-photoshop-info-format v0.0.0-20200609050348-3db9b63b202c // indirect
	github.com/dsoprea/go-utility/v2 v2.0.0-20200717064901-2fccff4aa15e // indirect
	github.com/go-errors/errors v1.1.1 // indirect
	github.com/go-xmlfmt/xmlfmt v0.0.0-20191208150333-d5b6f63a941b // indirect
	github.com/golang/geo v0.0.0-20200319012246-673a6f80352d // indirect
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
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
