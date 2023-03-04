package db

import (
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type DBController struct {
	log logrus.FieldLogger
	db  *mongo.Database
}

func NewDBController(log logrus.FieldLogger, db *mongo.Database) *DBController {
	return &DBController{
		log: log,
		db:  db,
	}
}
