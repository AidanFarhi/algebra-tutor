# syntax=docker/dockerfile:1

FROM golang:1.25.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

EXPOSE 8090

RUN CGO_ENABLED=0 GOOS=linux go build -o /algtutor

CMD [ "/algtutor" ]