FROM alpine:latest

# RUN apk add --no-cache ca-certificates

ENV REDIS_ADDRESS=redis:6379
ENV MONGO_DB_URL=mongo:27017
ENV MONGO_DB_NAME=access
# ENV REDIS_STREAM_GROUP=access_group
# ENV CHANNEL=access_channel
ENV HTTPPORT :5556
ENV DEFAULTPASS="defaultPass"
ENV TOKENSECRET="supersecret"

EXPOSE 5556
EXPOSE 9000
ADD main main

ENTRYPOINT [ "/main" ]
