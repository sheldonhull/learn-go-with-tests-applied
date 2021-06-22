module github.com/sheldonhull/learn-go-with-tests-applied/tools

// This is originally pulled from https://github.com/goyek/goyek and built upon
go 1.16

require (
	github.com/client9/misspell v0.3.4
	github.com/go-critic/go-critic v0.5.6
	github.com/golangci/golangci-lint v1.41.1
	github.com/goyek/goyek v0.5.0
	github.com/securego/gosec/v2 v2.8.0
	github.com/sqs/goreturns v0.0.0-20181028201513-538ac6014518
	golang.org/x/tools v0.1.3
	mvdan.cc/gofumpt v0.1.1
)

replace github.com/sheldonhull/learn-go-with-tests-applied => ../
