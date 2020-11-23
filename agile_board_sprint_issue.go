package jira

import (
	"fmt"
	"net/url"
)

// AgileBoardSprintIssue 返回指定Board的Sprint的Issue列表
func (cli *Client) AgileBoardSprintIssue(boardId, sprintId int, jql string) ([]*Issue, error) {
	var respInfo struct {
		CommonResponse
		Issues []*Issue
	}

	var query url.Values
	if jql != "" {
		query = url.Values{}
		query.Set("jql", jql)
	}

	err := cli.Request("GET", fmt.Sprintf("/rest/agile/1.0/board/%d/sprint/%d/issue", boardId, sprintId), query, &respInfo)
	if err != nil {
		return nil, err
	}

	return respInfo.Issues, nil
}
