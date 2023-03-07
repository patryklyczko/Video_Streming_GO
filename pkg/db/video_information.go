package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type VideoInformation struct {
	VideoName    string     `json:"video_name" bson:"video_name"`
	WatchedTimes int32      `json:"watched_times" bson:"watched_times"`
	Likes        int32      `json:"likes" bson:"likes"`
	LastWatched  *time.Time `json:"last_watched" bson:"last_watched"`
}

func (d *DBController) VideoInfo(name string) (*VideoInformation, error) {
	var videoInformation *VideoInformation
	var err error
	collection := d.db.Collection("VideoInfo")
	filter := bson.M{"filename": name}

	if err = collection.FindOne(context.Background(), filter).Decode(&videoInformation); err != nil {
		return nil, err
	}
	return videoInformation, nil
}
