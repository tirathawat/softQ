package http

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Info string `json:"info"`
}

func MakeCall(url string) (*Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response Response
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
