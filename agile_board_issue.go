package jira

import (
	"net/url"
)

// AgileBoardIssue 返回指定Board的Issue列表
func (cli *Client) AgileBoardIssue(boardId string, query url.Values) ([]*Issue, error) {
	var respInfo struct {
		CommonResponse
		Issues []*Issue
	}
	err := cli.Request("GET", "/rest/agile/1.0/board/"+boardId+"/issue", query, &respInfo)
	if err != nil {
		return nil, err
	}

	return respInfo.Issues, nil
}
