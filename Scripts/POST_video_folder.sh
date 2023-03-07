#!/bin/sh

# Example ./POST_video_folder.sh 
URL="http://localhost:8000/video/folder/manual"

data='{"path":"videos/fish.mp4",
    "name":"cat"}'

curl -X POST -d "$data" "$URL"  