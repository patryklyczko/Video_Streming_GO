package db

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VideoRequestFolder struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

type VideoRequest struct {
	Filename string `json:"filename"`
}

type VideoInfo struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Length     int64              `json:"length" bson:"length"`
	ChunkSize  int64              `json:"chunk_size" bson:"chunkSize"`
	UploadDate *time.Time         `json:"upload_data" bson:"uploadData"`
	FileName   string             `json:"filename" bson:"filename"`
}

type VideoFilter struct {
	Filename string `json:"filename" bson:"filename"`
}

func (d *DBController) PostVideoFolder(video VideoRequestFolder) error {
	collection := d.db.Collection("VideoInfo")
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

	videoInfo := VideoInformation{
		VideoName:    video.Name,
		WatchedTimes: 0,
		Likes:        0,
		LastWatched:  time.Now(),
	}
	_, err = collection.InsertOne(context.Background(), videoInfo)
	if err != nil {
		return err
	}

	return nil
}

func (d *DBController) WatchVideo(video VideoRequest) (*gridfs.DownloadStream, error) {
	collection := d.db.Collection("VideoInfo")
	var stream *gridfs.DownloadStream
	var err error
	if stream, err = d.bucket.OpenDownloadStreamByName("fish"); err != nil {
		return nil, err
	}

	filter := bson.M{"video_name": video.Filename}
	update := bson.M{"$inc": bson.M{"watched_times": 1}, "$set": bson.M{"last_watched": time.Now()}}
	opts := options.Update().SetUpsert(false)
	_, err = collection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
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

func (d *DBController) Videos(videoFilter VideoFilter) ([]VideoInfo, error) {
	var videos []VideoInfo

	filter := bson.M{"filename": bson.M{"$regex": ".*" + videoFilter.Filename + ".*", "$options": "i"}}
	cursor, err := d.bucket.Find(filter)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := cursor.Close(context.TODO()); err != nil {
			return
		}
	}()
	if err = cursor.All(context.TODO(), &videos); err != nil {
		log.Fatal(err)
	}
	return videos, nil
}

func (d *DBController) DeleteVideo(name string) error {
	collection := d.db.Collection("VideoInfo")
	var videos []VideoInfo
	var err error

	videoFilter := VideoFilter{
		Filename: name,
	}
	if videos, err = d.Videos(videoFilter); err != nil {
		return err
	}
	for _, video := range videos {
		if err := d.bucket.Delete(video.FileName); err != nil {
			return err
		}
	}

	filter := bson.M{"filename": name}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
