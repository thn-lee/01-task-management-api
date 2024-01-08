# all: go-build \
# 	docker-build \
# 	docker-save \
# 	docker-clean

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

docker-clean:
	docker image prune -f

go-mod tidy:
	go mod tidy
#  go get -v -u && go mod tidy

go-test:
	go test -v ./... -v -cover

go-run run r:
	go run ./cmd/server/
