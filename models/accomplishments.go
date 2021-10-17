package models

type Accomplishments struct {
	Languages []Language `yaml:"languages"`
	Projects  []Project  `yaml:"projects"`
}
