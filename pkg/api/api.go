package api

import (
	"github.com/patryklyczko/Video_Streming_GO/pkg/db"
	"github.com/sirupsen/logrus"
)

type InstanceAPI struct {
	log          logrus.FieldLogger
	dbController *db.DBController
}

func NewInstanceAPI(log logrus.FieldLogger, dbController *db.DBController) *InstanceAPI {
	return &InstanceAPI{
		log:          log,
		dbController: dbController,
	}
}
