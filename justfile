#!/usr/bin/env just --justfile

default:
    @just --list | grep -v default

tidy:
    go mod tidy

test PKG="./..." *ARGS="":
    go test -v -race -failfast -count 1 -coverprofile cover.out {{ PKG }} {{ ARGS }}

alias benchmark := bench

bench PKG="./..." *ARGS="":
    go test -v -count 1 -run xx -bench . {{ PKG }} {{ ARGS }}

lint PKG="./...":
    golangci-lint run --new=false {{ PKG }}

push *ARGS: test lint
    git push origin
