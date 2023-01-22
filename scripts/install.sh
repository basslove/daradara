#!/bin/bash -eu

go mod download

echo "run: go mod tidy"
go mod tidy
