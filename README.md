# go-server

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
   g
   ```

   
