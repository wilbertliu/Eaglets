package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/wilbertliu/Eaglets/credentials"
)

func main() {
	// Setup the Twitter client.
	anaconda.SetConsumerKey(credentials.ConsumerKey)
	anaconda.SetConsumerSecret(credentials.ConsumerSecret)
	twitterAPI := anaconda.NewTwitterApi(credentials.AccessToken, credentials.AccessTokenSecret)

	// Read the input from user.
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	// If there's error, print the error and stop the execution.
	if err != nil {
		fmt.Println(err)
		return
	}

	// Otherwise post status update on Twitter.
	twitterAPI.PostTweet(text, nil)
}
