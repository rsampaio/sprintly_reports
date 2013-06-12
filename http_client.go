package sprintly_reports

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Client struct {
	Username string
	Password string
	Client *http.Client
	
}

type Product struct {
	Name string
	Id int64
}

type Item struct {
	Who string
	What string
	Why string
	Tags []string
	Assigned_to string
}

func NewClient(user, pass string) (*Client) {
	c := new(Client)
	c.Username = user
	c.Password = pass
	c.Client = &http.Client{}
	return c
}

func (h *Client) HTTPRequest(method string, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println("request failed")
	}
	
	req.SetBasicAuth(h.Username, h.Password)
	resp, err := h.Client.Do(req)

	if err != nil {
		return nil, err
	}
	
	return resp, nil
}

func (h *Client) Products() ([]Product, error) {
	resp, _ := h.HTTPRequest("GET", "https://sprint.ly/api/products.json")
	body, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		return nil, err
	}

	var products []Product

	err = json.Unmarshal(body, &products)

	if err != nil {
		fmt.Println("json", err)
	}
	return products, nil
}

func (h *Client) Items(product_id int64) ([]Item, error) {
	var items []Item

	resp, err := h.HTTPRequest("GET", fmt.Sprintf("https://sprint.ly/api/products/%d/items.json?type=story", product_id))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	err = json.Unmarshal(body, &items)
	if err != nil {
		return nil, err
	}

	return items, nil;
}

