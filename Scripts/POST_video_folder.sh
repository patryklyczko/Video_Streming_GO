#!/bin/sh

URL="http://localhost:8000/video/folder/manual"

data='{"path":"videos/fish.mp4"}'

curl -X POST -d "$data" "$URL"  