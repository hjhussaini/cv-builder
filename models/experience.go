package models

type Experience struct {
	Tittle         string   `yaml:"tittle"`
	EmploymentType string   `yaml:"employment_type"`
	Company        string   `yaml:"company"`
	Location       string   `yaml:"location"`
	StartDate      string   `yaml:"start_date"`
	EndDate        string   `yaml:"end_date"`
	Description    []string `yaml:"description"`
}

func (experience *Experience) Date() string {
	return experience.StartDate + " - " + experience.EndDate
}
