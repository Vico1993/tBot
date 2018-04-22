package main

import (
	"fmt"

	"github.com/Vico1993/tbot/tbot"
	"github.com/sirupsen/logrus"
)

func main() {
	// // Let's configuration about twitter
	// viper.SetConfigName("config")
	// viper.AddConfigPath(".")
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	panic(fmt.Errorf("Fatal error config file: %s", err))
	// }

	// anaconda.SetConsumerKey(viper.GetString("twitter.consumerKey"))
	// anaconda.SetConsumerSecret(viper.GetString("twitter.consumerSecret"))
	// api := anaconda.NewTwitterApi(viper.GetString("twitter.accessToken"), viper.GetString("twitter.accessTokenSecret"))

	// log := &logger{logrus.New()}
	// api.SetLogger(log)

	// stream := api.PublicStreamFilter(url.Values{
	// 	"track":    []string{"vancouver"},
	// 	"location": []string{"49.246292, -123.116226."},
	// })

	// defer stream.Stop()

	// for v := range stream.C {
	// 	t, ok := v.(anaconda.Tweet)
	// 	if !ok {
	// 		log.Errorf("Receive unexpected value of type %T", v)
	// 		continue
	// 	}

	// 	tweetDesc := t.Text
	// 	tweetCount := t.RetweetCount
	// 	tweetAuthor := t.User.ScreenName

	// 	if t.RetweetedStatus != nil {
	// 		tweetDesc = t.RetweetedStatus.Text
	// 		tweetCount = t.RetweetedStatus.RetweetCount
	// 		tweetAuthor = t.RetweetedStatus.User.ScreenName
	// 	}

	// 	log.Infof("(@%s) - %s - %d \n", tweetAuthor, tweetDesc, tweetCount)
	// }

	tbot.GetTopTrends()

	fmt.Println("Build Works")
}

type logger struct {
	*logrus.Logger
}

func (log *logger) Critical(args ...interface{})                 { log.Error(args...) }
func (log *logger) Criticalf(format string, args ...interface{}) { log.Errorf(format, args...) }
func (log *logger) Notice(args ...interface{})                   { log.Info(args...) }
func (log *logger) Noticef(format string, args ...interface{})   { log.Infof(format, args...) }
