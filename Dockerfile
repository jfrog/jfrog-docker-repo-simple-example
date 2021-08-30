FROM gopherhut.jfrog.io/docker-virtual/golang:1.12.7-alpine3.9

MAINTAINER vignesh

RUN mkdir /app

ADD . /app
WORKDIR /app

RUN apk add git
RUN go get -u github.com/gorilla/mux

# RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...
CMD ["go","run","main.go"]

CMD printf "\nCONGRATULATIONS!!!\n\nYou have just set up your first Docker repository with the new JFrog Platform!\n\n"

