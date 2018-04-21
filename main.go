package main

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// Let's configuration about twitter
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	anaconda.SetConsumerKey(viper.GetString("twitter.consumerKey"))
	anaconda.SetConsumerSecret(viper.GetString("twitter.consumerSecret"))
	api := anaconda.NewTwitterApi(viper.GetString("twitter.accessToken"), viper.GetString("twitter.accessTokenSecret"))

	log := &logger{logrus.New()}
	api.SetLogger(log)

	stream := api.PublicStreamFilter(url.Values{
		"track": []string{"vancouver"},
	})

	defer stream.Stop()

	for v := range stream.C {
		t, ok := v.(anaconda.Tweet)
		if !ok {
			log.Errorf("Receive unexpected value of type %T", v)
			continue
		}

		log.Infof("%s\n", t.Text)
	}

	fmt.Println("Build Works")
}

type logger struct {
	*logrus.Logger
}

func (log *logger) Critical(args ...interface{})                 { log.Error(args...) }
func (log *logger) Criticalf(format string, args ...interface{}) { log.Errorf(format, args...) }
func (log *logger) Notice(args ...interface{})                   { log.Info(args...) }
func (log *logger) Noticef(format string, args ...interface{})   { log.Infof(format, args...) }
