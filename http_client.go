package sprintly_reports

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

type Client struct {
	username string
	password string
	client *http.Client
}

func NewClient(user, pass string) (*Client) {
	c := new(Client)
	c.username = user
	c.password = pass
	c.client = &http.Client{}
	return c
}

func (h *Client) Products() (string, error) {
	req, err := http.NewRequest("GET", "https://sprint.ly/api/products.json", nil)

	if err != nil {
		fmt.Println("request failed")
	}

	req.SetBasicAuth(h.username, h.password)
	resp, err := h.client.Do(req)

	if err != nil {
		return "failed", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		return "", err
	}
	return string(body), nil
}
