package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

const (
	targetTweetStatusID = 1085074572750155776
	until               = "2019-01-18"
	saveFilePath        = "./retweeted_users.csv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	config := oauth1.NewConfig(os.Getenv("CONSUMER_API_KEY"), os.Getenv("CONSUMER_API_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	savefile, err := os.Create(saveFilePath)
	if err != nil {
		panic(err)
	}
	defer savefile.Close()

	savefile.Write(([]byte)("User ID, Link, User Name, RetweetedAt, RetweetID\n"))

	targetTweet, _, _ := client.Statuses.Show(targetTweetStatusID, nil)

	tweetText := string([]rune(targetTweet.Text)[:61])
	fmt.Println(tweetText)
	var maxID int64 = 1111111111111111111

	for {
		search, _, _ := client.Search.Tweets(&twitter.SearchTweetParams{
			Count:   100,
			MaxID:   maxID,
			Query:   tweetText,
			SinceID: targetTweetStatusID,
			Until:   until,
		})

		if len(search.Statuses) == 0 {
			break
		}

		tmpLines := ""
		for _, status := range search.Statuses {
			if status.ID == maxID {
				continue
			}
			link := "https://twitter.com/" + status.User.ScreenName
			tmpLines += fmt.Sprintf("%s, %s, %s, %s, %d\n",
				status.User.ScreenName,
				link,
				status.User.Name,
				status.CreatedAt,
				status.ID,
			)
			fmt.Println(tmpLines)
			maxID = status.ID
		}
		if tmpLines == "" {
			break
		}
		savefile.Write(([]byte)(tmpLines))

	}
}
