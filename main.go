package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/wilbertliu/Eaglets/credentials"
)

func readTweets() ([]string, error) {
	// Make slice to store tweets.
	tweets := make([]string, 0)

	// Read tweets from user.
	reader := bufio.NewReader(os.Stdin)
	isStillReadingTweets := true
	for isStillReadingTweets {
		text, err := reader.ReadString('\n')

		// Return error immediately if any.
		if err != nil {
			return nil, err
		}

		// Append tweet to tweets slice if tweet is not empty.
		// Otherwise, stop reading tweets.
		if text != "\n" {
			tweets = append(tweets, text)
		} else {
			isStillReadingTweets = false
		}

		// Print newline to ease user distinguishes each tweet.
		fmt.Println()
	}

	return tweets, nil
}

func postTweets(tweets []string) error {
	// Setup the Twitter client.
	anaconda.SetConsumerKey(credentials.ConsumerKey)
	anaconda.SetConsumerSecret(credentials.ConsumerSecret)
	twitterAPI := anaconda.NewTwitterApi(credentials.AccessToken, credentials.AccessTokenSecret)

	// Post tweets.
	var postedTweet *anaconda.Tweet
	for _, tweet := range tweets {
		// Set the additional parameters.
		v := url.Values{}
		v.Set("trim_user", "true")
		v.Set("enable_dm_commands", "false")

		// When there's posted tweet, set the reply status id.
		if postedTweet != nil {
			v.Set("in_reply_to_status_id", postedTweet.IdStr)
		}

		// Post tweets one by one.
		postedTweetInLoop, err := twitterAPI.PostTweet(tweet, v)

		// Return error immediately if any.
		if err != nil {
			return err
		}

		// Store posted tweet object for further usage.
		postedTweet = &postedTweetInLoop
	}

	return nil
}

func main() {
	tweets, readingErr := readTweets()

	// If there's error, print and stop execution.
	if readingErr != nil {
		fmt.Println(readingErr)
		return
	}

	postingErr := postTweets(tweets)

	// If there's error, print and stop execution.
	if postingErr != nil {
		fmt.Println(postingErr)
		return
	}
}
