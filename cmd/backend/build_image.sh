#!/bin/sh

source_path=./cmd/backend
go_file=main.go
image_name=fireworks
build_output=fireworks
version=latest

CGO_ENABLED=0 GOOS="linux" GOARCH="amd64" go build -o $source_path/$build_output $source_path/$go_file

upx $source_path/$build_output

docker rmi -f $image_name:$version
docker build -f $source_path/Dockerfile -t dollarkiller/$image_name:$version  .
rm $source_path/$build_output

#docker save -o $image_name-$version.tar $image_name:$version
