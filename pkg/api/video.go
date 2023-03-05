package api

import "go.mongodb.org/mongo-driver/mongo/gridfs"

func (a *InstanceAPI) PostVideoFolder(url string) error {
	return a.dbController.PostVideoFolder(url)
}

func (a *InstanceAPI) WatchVideo(name string) (*gridfs.DownloadStream, error) {
	return a.dbController.WatchVideo(name)
}

func (a *InstanceAPI) AddVideo(file []byte, name string) error {
	return a.dbController.AddVideo(file, name)
}
