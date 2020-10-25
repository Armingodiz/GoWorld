package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
	"os"
)

// Credentials stores all of our access/consumer tokens
// and secret keys needed for authentication against
// the twitter REST API.

type Credentials struct {
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
}

func getCredentials() Credentials {
	var credentials Credentials
	jsonFile, err := os.Open("credentials.json")
	if err != nil {
		log.Fatalln(err)
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(jsonFile)
	err2 := json.Unmarshal(buf.Bytes(), &credentials)
	if err2 != nil {
		log.Fatalln(err2)
	}
	return credentials
}

func main() {
	UserInfo := getCredentials()
	client, err := getClient(&UserInfo)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}
	var input int
	for true {
		fmt.Println("1) tweet \n2) search \n3) show time line \n4) show followers \n5) show direct massages ")
		fmt.Scan(&input)
		switch input {
		case 1:
			sendTweet(client)
		case 2:
			search(client)
		case 3:
			timeLine(client)
		case 4:
			showFollowers(client)
		case 5:
			getDirects(client)
		}
	}
}

// getClient is a helper function that will return a twitter client
// that we can subsequently use to send tweets, or to stream new tweets
// this will take in a pointer to a Credential struct which will contain
// everything needed to authenticate and return a pointer to a twitter Client
// or an error

func getClient(creds *Credentials) (*twitter.Client, error) {
	// Pass in your consumer key (API Key) and your Consumer Secret (API Secret)
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	// Pass in your Access Token and your Access Token Secret
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	// we can retrieve the user and verify if the credentials
	// we have used successfully allow us to log in!
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}
	fmt.Println("user  " + user.ScreenName + " is connected ")
	return client, nil
}

func sendTweet(client *twitter.Client) {
	fmt.Println("ENTER THE TWEET : ")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	tweet, resp, err := client.Statuses.Update(input, &twitter.StatusUpdateParams{})
	if err != nil {
		fmt.Println(resp)
		log.Println(err)
	} else {
		fmt.Println(tweet.Text + "  TWEETED SUCCESSFULLY ! ")
	}
}

func search(client *twitter.Client) {
	fmt.Println("1 ) SEARCH AND SHOW \n2 ) SEARCH AND RETWEET \n3 ) SEARCH AND FOLLOW ")
	var input int
	fmt.Scan(&input)
	fmt.Println("ENTER THE TEXT YOU ARE LOOKING FOR : ")
	inputReader := bufio.NewReader(os.Stdin)
	txt, _ := inputReader.ReadString('\n')
	fmt.Println("ENTER COUNT OF RESULTS YOU WANT : ")
	var res int
	fmt.Scan(&res)
	search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: txt,
		Count: res,
	})
	if err != nil {
		fmt.Println(resp)
		log.Print(err)
	}
	for i, txt := range search.Statuses {
		switch input {
		case 1:
			fmt.Println("******************************************* tweet number " + string(i))
			fmt.Println(txt.Text)
		case 2:
			client.Statuses.Retweet(txt.ID, &twitter.StatusRetweetParams{})
		case 3:
			client.Friendships.Create(&twitter.FriendshipCreateParams{
				ScreenName: txt.User.ScreenName,
				UserID:     txt.User.ID,
				Follow:     nil,
			})
		}
	}
}

func timeLine(client *twitter.Client) {
	tweets, res, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{Count: 500})
	if err == nil {
		for _, tweet := range tweets {
			fmt.Println("*******************")
			fmt.Println(tweet.Text)
			fmt.Println("*******************")
		}
	} else {
		fmt.Println(res)
		log.Fatalln(err)
	}
}
func showFollowers(client *twitter.Client) {
	fmt.Println("ENTER COUNT OF RECENT FOLLOWERS  : ")
	var num int
	fmt.Scan(&num)
	followers, resp, err := client.Followers.List(&twitter.FollowerListParams{Count: num})
	if err == nil {
		for _, us := range followers.Users {
			fmt.Println("################## FOLLOWER " )
			fmt.Println(us.Name)
		}
	} else {
		fmt.Println(resp)
		log.Fatalln(err)
	}
}

func getDirects(client *twitter.Client) {
	directs, res, err := client.DirectMessages.EventsList(&twitter.DirectMessageEventsListParams{Count: 100})
	if err == nil {
		for _, direct := range directs.Events {
			fmt.Println("########################## NEW MESSAGE : ")
			fmt.Println(direct.Message.Data.Text)
		}
	} else {
		fmt.Println(res)
		log.Fatalln(err)
	}
}
