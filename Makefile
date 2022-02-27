check_install:
	which swagger || GO111MODULE=off go get -d -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	GO111MODULE=off swagger generate spec -t=./cmd/api/server -o ./api/swagger/swagger.yaml --scan-models