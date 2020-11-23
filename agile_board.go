package jira

import (
	"fmt"
	"net/url"
)

type Board struct {
	Id   int
	Self string
	Name string
	Type string
}

// AgileBoard 返回Board列表
func (cli *Client) AgileBoard(boardType, name, projectKeyOrId string) ([]*Board, error) {
	if boardType != "scrum" && boardType != "kanban" {
		return nil, fmt.Errorf("Invalid board type %s", boardType)
	}

	var respInfo struct {
		CommonResponse
		Values []*Board
	}

	query := url.Values{
		"type": {boardType},
	}

	if name != "" {
		query.Set("name", name)
	}

	if projectKeyOrId != "" {
		query.Set("projectKeyOrId", projectKeyOrId)
	}

	err := cli.Request("GET", "/rest/agile/1.0/board", query, &respInfo)
	if err != nil {
		return nil, err
	}

	return respInfo.Values, nil
}
