package db

import (
	"io/ioutil"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type VideoRequestFolder struct {
	Path string `json:"path"`
}

type Video struct {
	Name    string             `json:"name"`
	VideoID primitive.ObjectID `json:"video_id"`
}

type FileChunk struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	FilesID primitive.ObjectID `bson:"files_id,omitempty"`
	N       int32              `bson:"n,omitempty"`
	Data    []byte             `bson:"data,omitempty"`
}

func (d *DBController) PostVideoFolder(path string) error {
	videoBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	uploadStream, err := d.bucket.OpenUploadStream(
		"fish.mp4",
	)
	if _, err = uploadStream.Write(videoBytes); err != nil {
		return nil
	}

	if err := uploadStream.Close(); err != nil {
		return err
	}

	return nil
}

func (d *DBController) WatchVideo(name string) (*gridfs.DownloadStream, error) {
	var stream *gridfs.DownloadStream
	var err error
	if stream, err = d.bucket.OpenDownloadStreamByName("fish.mp4"); err != nil {
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
