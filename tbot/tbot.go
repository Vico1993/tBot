package tbot

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var api *anaconda.TwitterApi
var log = &logger{logrus.New()}

func init() {
	// Let's configuration about twitter
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	anaconda.SetConsumerKey(viper.GetString("twitter.consumerKey"))
	anaconda.SetConsumerSecret(viper.GetString("twitter.consumerSecret"))
	api = anaconda.NewTwitterApi(viper.GetString("twitter.accessToken"), viper.GetString("twitter.accessTokenSecret"))
	api.SetLogger(log)

}

// GetSpecificStream return a direct stream from a specifique element.
func GetSpecificStream(keyword string) {

	stream := api.PublicStreamFilter(url.Values{
		"track": []string{keyword},
	})

	defer stream.Stop()

	for v := range stream.C {
		t, ok := v.(anaconda.Tweet)
		if !ok {
			log.Errorf("Receive unexpected value of type %T", v)
			continue
		}

		tweetDesc := t.Text
		tweetCount := t.RetweetCount
		tweetAuthor := t.User.ScreenName

		if t.RetweetedStatus != nil {
			tweetDesc = t.RetweetedStatus.Text
			tweetCount = t.RetweetedStatus.RetweetCount
			tweetAuthor = t.RetweetedStatus.User.ScreenName
		}

		log.Infof("(@%s) - %s - %d \n", tweetAuthor, tweetDesc, tweetCount)
	}
}

// GetTopTrends for a specific location
func GetTopTrends(lat float64, long float64) {
	// Get WHOEID from yahoo of a specific location
	location, err := api.GetTrendsClosestLocations(lat, long, url.Values{})
	if err != nil || len(location) > 1 || len(location) <= 0 {
		panic(fmt.Errorf("Error receiving the Location: %s", err))
	}

	woeid := location[0].Woeid

	// Get Trends for this location
	trendResponses, err := api.GetTrendsByPlace(int64(woeid), url.Values{})
	if err != nil {
		panic(fmt.Errorf("Error receiving trends: %s", err))
	}

	for _, t := range trendResponses.Trends {
		if t.PromotedContent != "" {
			continue
		}
		log.Info(t.Name + " - " + t.Url)
	}
}

type logger struct {
	*logrus.Logger
}

func (log *logger) Critical(args ...interface{})                 { log.Error(args...) }
func (log *logger) Criticalf(format string, args ...interface{}) { log.Errorf(format, args...) }
func (log *logger) Notice(args ...interface{})                   { log.Info(args...) }
func (log *logger) Noticef(format string, args ...interface{})   { log.Infof(format, args...) }
