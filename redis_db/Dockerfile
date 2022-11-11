#use an existing docker image as a base 
FROM alpine

#Download and install a dependency 
RUN apk add --update redis 

EXPOSE 6379

COPY redis.conf .
#startup command 
CMD [ "redis-server", "./redis.conf" ]
