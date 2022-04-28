module photo-manager-cli/action

go 1.18

require (
	photo-manager-cli/album v0.0.0
	photo-manager-cli/config v0.0.0
	photo-manager-cli/metadata v0.0.0
	photo-manager-cli/uploader v0.0.0
)

replace (
	photo-manager-cli/album v0.0.0 => ../album
	photo-manager-cli/config v0.0.0 => ../config
	photo-manager-cli/metadata v0.0.0 => ../metadata
	photo-manager-cli/uploader v0.0.0 => ../uploader
)
