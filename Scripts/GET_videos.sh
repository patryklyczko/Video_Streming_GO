#!/bin/bash

# Example ./GET_videos.sh "filename=fish"
URL="http://localhost:8000/video/name"

QUERY_PARAMS=${1:-}

echo "Query params: $QUERY_PARAMS"
if [ ! -z "$QUERY_PARAMS" ]; then
  URL="$URL?$QUERY_PARAMS"
fi

RESPONSE=$(curl -X GET "$URL")
echo "$RESPONSE" | jq
