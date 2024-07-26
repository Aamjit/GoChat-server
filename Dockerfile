FROM golang:1.22.5
# RUN apk add build-base
# RUN apk add --no-cache git
# Do a system update
# RUN apk update
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod download
ENV APP_ENV=PRD
RUN go build -o main
CMD ["/app/main"]