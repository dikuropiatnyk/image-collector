# Use the official GoLang base image
FROM golang:1.17.3

LABEL maintainer="Danylo Kuropiatnyk <dikuropiatnyk@spscommerce.com>"

ARG APP_PATH=/app

# Create APP folder:
RUN mkdir -p ${APP_PATH}

WORKDIR ${APP_PATH}

# Copy the Go module files and download dependencies
COPY go.mod ${APP_PATH}/go.mod
COPY go.sum ${APP_PATH}/go.sum
RUN go mod download

# Copy the Go source code into the container
COPY main.go ${APP_PATH}/main.go
COPY ./models ${APP_PATH}/models
COPY ./services ${APP_PATH}/services
COPY ./utils ${APP_PATH}/utils

# Build the Go application
RUN go build -o app

# Run the Go application
CMD ["./app"]
