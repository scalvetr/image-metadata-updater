module photo-manager-cli/album

go 1.23

require (
	photo-manager-cli/config v0.0.0
)

replace (
	photo-manager-cli/config v0.0.0 => ../config
)
