package api

import (
	"github.com/patryklyczko/Video_Streming_GO/pkg/db"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func (a *InstanceAPI) PostVideoFolder(video db.VideoRequestFolder) error {
	return a.dbController.PostVideoFolder(video)
}

func (a *InstanceAPI) WatchVideo(video db.VideoRequest) (*gridfs.DownloadStream, error) {
	return a.dbController.WatchVideo(video)
}

func (a *InstanceAPI) AddVideo(file []byte, name string) error {
	return a.dbController.AddVideo(file, name)
}
