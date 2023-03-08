FROM golang:1.19.4-alpine
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

RUN apk add --update && apk add git && apk add postgresql-client

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

COPY . .
EXPOSE 8080

RUN go install github.com/cosmtrek/air@latest
CMD ["air"]