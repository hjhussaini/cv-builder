package models

type Country struct {
	Name     string `yaml:"name"`
	Location string `yaml:"location,omitempty"`
}
