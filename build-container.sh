#!/bin/sh
docker start go-builder
if [ $# -eq 1 ]; then
  docker exec -it go-builder sh -c "cd /app && sh build.sh $1"
  docker-compose up -d --build $1
else
  docker exec -it go-builder sh -c "cd /app && sh build.sh"
  docker-compose up -d --build
fi
