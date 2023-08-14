#!/bin/bash
set -ex

if ! command -v gh >/dev/null; then
  echo must install gh
  echo brew install gh
  exit 1
fi

TAG=$(grep 'VERSION =' ls-go.go | grep -oE '\d+\.\d+.\d+')

git tag --force latest
git tag -a "v$TAG" -m "release v$TAG"

git push origin master --tags --force

./compile

gh release create "v$TAG"
gh release upload "v$TAG" ls-go-*
