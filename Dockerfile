FROM golang:1.17.7-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /web-service-gin
EXPOSE 8080
CMD [ "/web-service-gin" ]
