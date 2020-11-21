package jira

type User struct {
	Self         string
	Name         string
	Key          string
	EmailAddress string
	AvatarUrls   map[string]string
	DisplayName  string
	Active       bool
	TimeZone     string
}
