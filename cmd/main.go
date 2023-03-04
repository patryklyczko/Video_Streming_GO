package main

import (
	"context"
	"fmt"

	"github.com/patryklyczko/Video_Streming_GO/pkg/api"
	"github.com/patryklyczko/Video_Streming_GO/pkg/db"
	"github.com/patryklyczko/Video_Streming_GO/pkg/http"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	confOptMongoPassword = 
	confOptMongoUser     = "Video_GO"
	confOptMongoDatabase = "Video_GO"
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	server := createServerFromConfig(logger, ":8000")
	server.Run()
}

func createServerFromConfig(logger *logrus.Logger, bind string) *http.HTTPInstanceAPI {

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	mongoURI := fmt.Sprintf(
		"mongodb+srv://%s:%s@videogo.p3ubuzc.mongodb.net/?retryWrites=true&w=majority&ssl=true",
		confOptMongoUser,
		confOptMongoPassword)

	clientOptions := options.Client().
		ApplyURI(mongoURI).
		SetServerAPIOptions(serverAPIOptions)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		logger.WithError(err).Fatal("could not instatiate ")
	}

	dbController := db.NewDBController(
		logger.WithField("component", "db"),
		client.Database(viper.GetString(confOptMongoDatabase)),
	)
	instanceAPI := api.NewInstanceAPI(
		logger.WithField("component", "api"),
		dbController,
	)

	return http.NewHTTPInstanceAPI(bind, logger.WithField("component", "http"), instanceAPI)
}
