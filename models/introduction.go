package models

type Introduction struct {
	Photo     string  `yaml:"photo,omitempty"`
	FirstName string  `yaml:"first_name"`
	LastName  string  `yaml:"last_name"`
	Headline  string  `yaml:"headline"`
	Country   Country `yaml:"country"`
}

func (introduction *Introduction) FullName() string {
	return introduction.FirstName + " " + introduction.LastName
}
