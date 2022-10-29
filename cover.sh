#!/bin/bash

go test $(go list ./... | grep -vE "(vendor)|(test)|(array$)|(mock)|(bench)|(checker)|(proto)|(interface)") -race -count=1 -coverprofile=coverage.out -args "postgres://postgres:qaz@localhost:5432/mydb"
go tool cover -func=coverage.out