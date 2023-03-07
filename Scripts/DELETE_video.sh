#!/bin/sh

URL="http://localhost:8000/video"

data='{"filename":"cat"}'

curl -X DELETE -d "$data" "$URL"
