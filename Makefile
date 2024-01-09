# all: go-build \
# 	docker-build \
# 	docker-save \
# 	docker-clean

install:
	brew install go

go-install:
	go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest

field-align f-al:
	fieldalignment -fix ./...

go-build:
	CGO_ENABLED=0 go build -v \
	-buildvcs=false \
	-installsuffix 'static' \
	-ldflags="-X 'main.version=$$(git rev-parse --short HEAD)' -X 'main.build=$$(date --iso-8601=seconds)'" \
	-o ./dist/server \
	./cmd/server

# docker-build:
# 	docker build -f ./cmd/server/Dockerfile \
# 	-t thn-lee/01-task-management-api \
# 	--pull \
# 	.

# docker-save:
# 	docker save thn-lee/01-task-management-api | gzip > dist/01-task-management-api.tar.gz

# docker-clean:
# 	docker image prune -f

go-mod tidy m: 
	go mod tidy --compat=1.18

go-test test t:
	go test -v ./... -v -cover

go-run run r:
	go run ./cmd/server/
