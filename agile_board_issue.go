package jira

import (
	"fmt"
	"net/url"
)

// AgileBoardIssue 返回指定Board的Issue列表
func (cli *Client) AgileBoardIssue(boardId int, jql string) ([]*Issue, error) {
	var respInfo struct {
		CommonResponse
		Issues []*Issue
	}

	var query url.Values
	if jql != "" {
		query = url.Values{}
		query.Set("jql", jql)
	}

	err := cli.Request("GET", fmt.Sprintf("/rest/agile/1.0/board/%d/issue", boardId), query, &respInfo)
	if err != nil {
		return nil, err
	}

	return respInfo.Issues, nil
}
