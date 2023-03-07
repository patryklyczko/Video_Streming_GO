package db

import (
	"io/ioutil"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type VideoRequestFolder struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

type VideoRequest struct {
	Name string `json:"name"`
}

type FileChunk struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	FilesID primitive.ObjectID `bson:"files_id,omitempty"`
	N       int32              `bson:"n,omitempty"`
	Data    []byte             `bson:"data,omitempty"`
}

type VideoFilter struct {
	Name   string `json:"name"`
	Chunks int32  `json:"chunks"`
	IsNew  bool   `json:"is_new"`
}

func (d *DBController) PostVideoFolder(video VideoRequestFolder) error {
	videoBytes, err := ioutil.ReadFile(video.Path)
	if err != nil {
		return err
	}

	uploadStream, err := d.bucket.OpenUploadStream(
		video.Name,
	)
	if _, err = uploadStream.Write(videoBytes); err != nil {
		return nil
	}

	if err := uploadStream.Close(); err != nil {
		return err
	}

	return nil
}

func (d *DBController) WatchVideo(video VideoRequest) (*gridfs.DownloadStream, error) {
	var stream *gridfs.DownloadStream
	var err error
	if stream, err = d.bucket.OpenDownloadStreamByName(video.Name); err != nil {
		return nil, err
	}

	return stream, nil
}

func (d *DBController) AddVideo(file []byte, name string) error {
	uploadStream, err := d.bucket.OpenUploadStream(
		name,
	)
	if _, err = uploadStream.Write(file); err != nil {
		return nil
	}

	if err := uploadStream.Close(); err != nil {
		return err
	}

	return nil
}
