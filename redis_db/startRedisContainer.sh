#!/bin/sh
docker build . -t gta4roy/simple-redis-db:latest 
docker run  -p 6379:6379 -d gta4roy/simple-redis-db:latest  