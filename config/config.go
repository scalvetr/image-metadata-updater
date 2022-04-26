package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type Action uint

const (
	UpdateDateFromMetadata Action = iota
	UpdateDate
)

var ActionToString = map[Action]string{
	UpdateDateFromMetadata: "UPDATE_DATE_FROM_METADATA",
	UpdateDate:             "UPDATE_DATE",
}

var ActionFromString = map[string]Action{
	"UPDATE_DATE_FROM_METADATA": UpdateDateFromMetadata,
	"UPDATE_DATE":               UpdateDate,
}

func (a Action) String() string {
	if s, ok := ActionToString[a]; ok {
		return s
	}
	return "unknown"
}

func (a Action) MarshalYAML() ([]byte, error) {
	if s, ok := ActionToString[a]; ok {
		return yaml.Marshal(s)
	}
	return nil, fmt.Errorf("unknown user type %d", a)
}

func (a *Action) UnmarshalYAML(text []byte) error {
	var s string
	if err := yaml.Unmarshal(text, &s); err != nil {
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
	Action Action `yaml:"action" `
	Path   string `yaml:"path"`
	Date   string `yaml:"date"`
}
