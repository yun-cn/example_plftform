#!/bin/bash

appname=$1

url="http://localhost:1358/?&appname=${appname}&url=http://127.0.0.1:9200&mode=edit"

echo "opening :$url"

docker run --rm --name "noonde-dejavu" -p 1358:1358 -d appbaseio/dejavu
open "$url"
