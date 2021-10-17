package models

type Contact struct {
	LinkedIn string `yaml:"linkedin,omitempty"`
	Email    string `yaml:"email"`
	Mobile   string `yaml:"mobile"`
	Skype    string `yaml:"skype,omitempty"`
	Twitter  string `yaml:"twitter,omitempty"`
}
