#!/bin/sh

set -ex

go build -gcflags=-m

echo ""

go test -bench BenchmarkSortString1K

echo ""

go test -bench BenchmarkSortString1K sort
