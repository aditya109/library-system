check_api_directory:
	[ -d api ] || mkdir -p api/swagger

check_install: check_api_directory
	which swagger || GO111MODULE=off go install github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	GO111MODULE=off swagger generate spec -t=./cmd/server -o ./api/swagger/swagger.yaml --scan-models