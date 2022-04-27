module image-metadata-updater/action

go 1.18

require (
	image-metadata-updater/album v0.0.0
	image-metadata-updater/config v0.0.0
	image-metadata-updater/metadata v0.0.0
	image-metadata-updater/uploader v0.0.0
)

replace (
	image-metadata-updater/album v0.0.0 => ../album
	image-metadata-updater/config v0.0.0 => ../config
	image-metadata-updater/metadata v0.0.0 => ../metadata
	image-metadata-updater/uploader v0.0.0 => ../uploader
)
