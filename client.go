package jira

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// ApiResponse 表示通用的JIRA响应
type CommonResponse struct {
	Expand     string
	StartAt    int
	MaxResults int
	Total      int
}

type Client struct {
	http.Client

	ServerAddr string
	Username   string
	Password   string
}

func (cli *Client) Request(method, path string, query url.Values, respInfo interface{}) error {
	if query != nil {
		path += "?" + query.Encode()
	}

	req, err := http.NewRequest(method, cli.ServerAddr+path, nil)
	req.SetBasicAuth(cli.Username, cli.Password)

	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&respInfo)
	if err != nil {
		return err
	}

	return nil
}
