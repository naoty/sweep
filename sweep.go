package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	pocket "github.com/motemen/go-pocket/api"
)

type Envs struct {
	ConsumerKey string
	AccessToken string
	Expiration  int
}

const (
	DefaultExpiration = 24
)

var helpMsg = `Usage: sweep

Configuration:
  POCKET_CONSUMER_KEY: A consumer key for your Pocket application
  POCKET_ACCESS_TOKEN: An access token for your Pocket application
  POCKET_EXPIRATION: An expiration hours after which items are deleted (Default: 24)
`

type Sweep struct{}

func (s *Sweep) Help() string {
	return helpMsg
}

func (s *Sweep) Run(args []string) int {
	envs, err := loadEnvs()

	if err != nil {
		fmt.Println(err)
		return 1
	}

	client := pocket.NewClient(envs.ConsumerKey, envs.AccessToken)

	option := &pocket.RetrieveOption{}
	result, err := client.Retrieve(option)

	if err != nil {
		fmt.Println(err)
		return 1
	}

	expired := selectExpired(result.List, envs.Expiration)

	actions := []*pocket.Action{}
	for _, item := range expired {
		fmt.Println(item.GivenURL)
		action := &pocket.Action{Action: "delete", ItemID: item.ItemID}
		actions = append(actions, action)
	}

	_, err = client.Modify(actions...)

	if err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}

func (s *Sweep) Synopsis() string {
	return "Delete unread items"
}

func loadEnvs() (envs Envs, err error) {
	consumerKey := os.Getenv("POCKET_CONSUMER_KEY")
	accessToken := os.Getenv("POCKET_ACCESS_TOKEN")
	expirationStr := os.Getenv("POCKET_EXPIRATION")

	if consumerKey == "" || accessToken == "" {
		err = errors.New("Consumer key and access token is required.")
	}

	var expiration int
	if expirationStr == "" {
		expiration = DefaultExpiration
	} else {
		expiration, err = strconv.Atoi(expirationStr)
	}

	envs = Envs{
		ConsumerKey: consumerKey,
		AccessToken: accessToken,
		Expiration:  expiration,
	}

	return
}

// The unit of expiration is hour.
func selectExpired(items map[string]pocket.Item, expiration int) (expired []pocket.Item) {
	ed := time.Duration(expiration)
	for _, item := range items {
		t := time.Time(item.TimeAdded)
		d := time.Since(t)
		if d > ed*time.Hour {
			expired = append(expired, item)
		}
	}
	return
}
