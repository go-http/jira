package jira

type StatusCategory struct {
	Self      string
	Id        int
	Key       string
	ColorName string
	Name      string
}

type Status struct {
	Self           string
	Description    string
	IconUrl        string
	Name           string
	Id             string
	StatusCategory StatusCategory
}
