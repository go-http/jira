package jira

import (
	"fmt"
	"net/url"
)

// Issue 返回Sprint的Issue列表
func (sprint *Sprint) Issue(jql string) ([]*Issue, error) {
	var respInfo struct {
		CommonResponse
		Issues []*Issue
	}

	var query url.Values
	if jql != "" {
		query = url.Values{}
		query.Set("jql", jql)
	}

	err := sprint.Request("GET", fmt.Sprintf("/rest/agile/1.0/board/%d/sprint/%d/issue", sprint.Board.Id, sprint.Id), query, &respInfo)
	if err != nil {
		return nil, err
	}

	for _, issue := range respInfo.Issues {
		issue.Client = sprint.Client
	}

	return respInfo.Issues, nil
}
