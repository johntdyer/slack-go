package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Domain string
	Token  string
}

type Message struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconUrl   string `json:"icon_url"`
	IconEmoji string `json:"icon_emoji"`
	Channel   string `json:"channel"`
}

type SlackError struct {
	Code int
	Body string
}

func (e *SlackError) Error() string {
	return fmt.Sprintf("SlackError: %d %s", e.Code, e.Body)
}

func NewClient(domain, token string) *Client {
	return &Client{domain, token}
}

func (c *Client) getUrl() string {
	return fmt.Sprintf("https://%s/services/hooks/incoming-webhook?token=%s", c.Domain, c.Token)
}

func (c *Client) SendMessage(msg *Message) error {

	body, _ := json.Marshal(msg)
	fmt.Println(string(body))
	buf := bytes.NewReader(body)

	http.NewRequest("POST", c.getUrl(), buf)
	resp, err := http.Post(c.getUrl(), "application/json", buf)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		t, _ := ioutil.ReadAll(resp.Body)
		return &SlackError{resp.StatusCode, string(t)}
	}

	return nil
}
