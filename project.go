package jira

type ProjectCategory struct {
	Self        string
	Id          string
	Name        string
	Description string
}

type Project struct {
	Self            string
	Id              string
	Key             string
	Name            string
	AvatarUrls      map[string]string
	ProjectCategory ProjectCategory
}
