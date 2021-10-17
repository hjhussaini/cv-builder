package models

type Sections struct {
	Introduction    *Introduction    `yaml:"introduction"`
	Contact         *Contact         `yaml:"contact"`
	About           []string         `yaml:"about"`
	Skills          []string         `yaml:"skills"`
	Technologies    []string         `yaml:"tools_technologies"`
	Experiences     []Experience     `yaml:"experience"`
	Educations      []Education      `yaml:"education"`
	Accomplishments *Accomplishments `yaml:"accomplishments"`
}
