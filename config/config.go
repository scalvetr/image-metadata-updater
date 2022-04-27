package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type Action uint

const (
	Undefined Action = iota
	UpdateDateFromMetadata
	UpdateMetadataDate
	UploadAlbums
)

var ActionFromString = map[string]Action{
	"UNDEFINED":                 Undefined,
	"UPDATE_DATE_FROM_METADATA": UpdateDateFromMetadata,
	"UPDATE_METADATA_DATE":      UpdateMetadataDate,
	"UPLOAD_ALBUMS":             UploadAlbums,
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
	Action    Action    `yaml:"action" `
	Path      string    `yaml:"path"`
	Date      string    `yaml:"date"`
	GoogleApi GoogleApi `yaml:"google_api"`
}

type GoogleApi struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	Token        string `yaml:"token"`
}
