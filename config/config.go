package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type Action uint

const (
	Undefined Action = iota
	UpdateDateFromMetadata
	UpdateMetadata
	UploadAlbums
	FixDateAlbums
)

var ActionFromString = map[string]Action{
	"UNDEFINED":                 Undefined,
	"UPDATE_DATE_FROM_METADATA": UpdateDateFromMetadata,
	"UPDATE_METADATA":           UpdateMetadata,
	"UPLOAD_ALBUMS":             UploadAlbums,
	"FIX_DATE_ALBUMS":           FixDateAlbums,
}

func (a *Action) UnmarshalYAML(value *yaml.Node) error {
	var s string
	if err := value.Decode(&s); err != nil {
		return err
	}
	var v Action
	var ok bool
	if v, ok = ActionFromString[s]; !ok {
		return fmt.Errorf("unknown user type %s", s)
	}
	*a = v
	return nil
}

type Config struct {
	Action                   Action               `yaml:"action" `
	Path                     string               `yaml:"path"`
	Regexp                   string               `yaml:"regexp"`
	UpdateMetadataDateConfig UpdateMetadataConfig `yaml:"update_metadata_config"`
	AlbumInfoConfig          AlbumInfoConfig      `yaml:"album_info_config"`
}

type AlbumInfoConfig struct {
	FolderRegexp     string `yaml:"folder_regexp"`
	AlbumNamePattern string `yaml:"album_name_pattern"`
}

func (c AlbumInfoConfig) GetFolderRegexp() string {
	if c.FolderRegexp == "" {
		return `(?<year>\d{4}) - (?<month>\d{2})(.*) - (?<name>.*)`
	}
	return c.FolderRegexp

}
func (c AlbumInfoConfig) GetAlbumNamePattern() string {
	if c.AlbumNamePattern == "" {
		return `{{printf "%04d" .Year}}-{{printf "%02d" .Month}} - {{.Name}}`
	}
	return c.AlbumNamePattern
}

type UpdateMetadataConfig struct {
	Date         string                            `yaml:"date"`
	Override     bool                              `yaml:"override"`
	DateReplaces []UpdateMetadataDateConfigReplace `yaml:"date_replaces"`
}
type UpdateMetadataDateConfigReplace struct {
	Day    string `yaml:"day"`
	NewDay string `yaml:"new_day"`
}
