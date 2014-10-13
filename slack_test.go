package slack

import (
	"flag"
	"os/user"
	"testing"
)

var (
	client *Client
)

func init() {
	client = &Client{}
	flag.StringVar(&client.Domain, "domain", "", "slack domain")
	flag.StringVar(&client.Token, "token", "", "slack token")
	flag.Parse()
	if client.Domain == "" || client.Token == "" {
		flag.PrintDefaults()
		panic("\n=================\nYou need to specify -domain and -token flags\n=================\n\n")
	}
}

func TestSendMessage(t *testing.T) {
	msg := &Message{}
	msg.Channel = "#docker"
	msg.Text = "Slack API Test from go"
	user, _ := user.Current()
	msg.Username = user.Username
	client.SendMessage(msg)
}
