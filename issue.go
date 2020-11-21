package jira

import (
	"encoding/json"
)

type IssueAttributes struct {
	Summary     string
	Description string
	Assignee    User
	Reporter    User
	Creator     User
	Flagged     bool
	WorkRatio   float32

	Project Project
	Status  Status

	Created string
	Updated string

	IssueType struct {
		Self        string
		Id          string
		Description string
		IconUrl     string
		Name        string
		Subtask     bool
		AvatarId    int
	}

	Watches struct {
		Self       string
		WatchCount int
		IsWatching bool
	}

	Priority struct {
		Self    string
		IconUrl string
		Name    string
		Id      string
	}

	//TODO: 完善下列字段的结构
	timetracking                  interface{}
	environment                   interface{}
	fixVersions                   interface{}
	subtasks                      interface{}
	votes                         interface{}
	duedate                       interface{}
	aggregatetimeestimate         interface{}
	progress                      interface{}
	timeoriginalestimate          interface{}
	lastViewed                    interface{}
	timespent                     interface{}
	resolution                    interface{}
	aggregateprogress             interface{}
	labels                        interface{}
	epic                          interface{}
	worklog                       interface{}
	updated                       interface{}
	aggregatetimeoriginalestimate interface{}
	sprint                        interface{}
	security                      interface{}
	timeestimate                  interface{}
	resolutiondate                interface{}
	attachment                    interface{}
	issuelinks                    interface{}
	aggregatetimespent            interface{}
	comment                       interface{}
}

type Issue struct {
	IssueAttributes

	Id     string
	Self   string
	Key    string
	Expand string
	Fields json.RawMessage
}

func (issue *Issue) ParseAttributes() error {
	attr := IssueAttributes{}
	err := json.Unmarshal(issue.Fields, &attr)
	if err != nil {
		return err
	}

	issue.IssueAttributes = attr
	issue.Fields = nil

	return nil
}
