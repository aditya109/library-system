# go-server
[![BCH compliance](https://bettercodehub.com/edge/badge/aditya109/library-system?branch=main)](https://bettercodehub.com/)

## How to run server ?

1. Install all the dependencies.

   ```shell
   go mod download
   ```

2. Run a `mongo-db` container.

   ```shell
   docker run --name mongo-db -p 27017:27017 -d mongo:latest
   ```

3. Run the `go` server.

   ```bash
   go run cmd/main.go
   ```

   ```bash
   docker build --tag daitya96/library-server .; docker rm --force server; docker run --name server --publish 8000:8000 daitya96/library-server
   ```
docker run -it --rm -p 8000:8000 daitya96/library-server /bin/ash