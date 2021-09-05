# starting from golang base image
FROM golang:alpine as builder

# enable go modules
ENV GO111MODULE=on

# install git
RUN apk update && apk add --no-cache git

# set current working directory
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

# download all dependencies
RUN go mod download

# copy source code
COPY . .

# Note here: CGO_ENABLED is disabled for cross system compilation
# It is also a common best practise.
# Build the application.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main .

# Finally our multi-stage to build a small image
# Start a new stage from scratch
FROM scratch

# Copy the Pre-built binary file
COPY --from=builder /app/bin/main .

# Run executable
CMD ["./main"]