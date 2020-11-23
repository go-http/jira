package jira

import (
	"fmt"
	"net/url"
	"strings"
)

type Sprint struct {
	*Board

	Id            int
	Self          string
	Name          string
	State         string
	StartDate     string
	EndDate       string
	CompleteDate  string
	OriginBoardId int
}

// Sprint 返回Board的Sprint列表
func (board *Board) Sprint(states ...string) ([]*Sprint, error) {
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

	err := board.Request("GET", fmt.Sprintf("/rest/agile/1.0/board/%d/sprint", board.Id), query, &respInfo)
	if err != nil {
		return nil, err
	}

	for _, sprint := range respInfo.Values {
		sprint.Board = board
	}

	return respInfo.Values, nil
}
