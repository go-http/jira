package jira

import (
	"fmt"
	"net/url"
)

// Issue 返回Board的Issue列表
func (board *Board) Issue(jql string) ([]*Issue, error) {
	var respInfo struct {
		CommonResponse
		Issues []*Issue
	}

	var query url.Values
	if jql != "" {
		query = url.Values{}
		query.Set("jql", jql)
	}

	err := board.Request("GET", fmt.Sprintf("/rest/agile/1.0/board/%d/issue", board.Id), query, &respInfo)
	if err != nil {
		return nil, err
	}

	return respInfo.Issues, nil
}
