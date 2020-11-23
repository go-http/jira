package jira

import (
	"fmt"
	"net/url"
	"strings"
)

type Sprint struct {
	Id            int
	Self          string
	Name          string
	State         string
	StartDate     string
	EndDate       string
	CompleteDate  string
	OriginBoardId int
}

// AgileBoardSprint 返回Board中的Sprint列表
func (cli *Client) AgileBoardSprint(boardId int, states ...string) ([]*Sprint, error) {
	var respInfo struct {
		CommonResponse
		Values []*Sprint
	}

	for _, state := range states {
		if state != "future" && state != "active" && state != "closed" {
			return nil, fmt.Errorf("Invalid sprint state %s", state)
		}
	}

	query := url.Values{
		"state": {strings.Join(states, ",")},
	}

	err := cli.Request("GET", fmt.Sprintf("/rest/agile/1.0/board/%d/sprint", boardId), query, &respInfo)
	if err != nil {
		return nil, err
	}

	return respInfo.Values, nil
}
