#!/bin/bash
set -ex

if [ -z $GITHUB_TOKEN ]; then
  echo must set GITHUB_TOKEN >&2
  exit 1
fi

TAG=0.1.0

git tag -a v$TAG -m "release v$TAG"

git push origin master --tags

github-release release \
  --user acarl005 \
  --repo ls-go \
  --tag v$TAG \
  --name v$TAG

GOOS=darwin GOARCH=amd64 go build -o ls-go-darwin-amd64
GOOS=linux GOARCH=amd64 go build -o ls-go-linux-amd64
GOOS=linux GOARCH=arm64 go build -o ls-go-linux-arm64

github-release upload \
  --user acarl005 \
  --repo ls-go \
  --tag v$TAG \
  --name ls-go-darwin-amd64 \
  --file ls-go-darwin-amd64

github-release upload \
  --user acarl005 \
  --repo ls-go \
  --tag v$TAG \
  --name ls-go-linux-amd64 \
  --file ls-go-linux-amd64

github-release upload \
  --user acarl005 \
  --repo ls-go \
  --tag v$TAG \
  --name ls-go-linux-arm64 \
  --file ls-go-linux-arm64
