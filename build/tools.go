// +build tools

package tools

// Manage tool dependencies via go.mod.
//
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
// https://github.com/golang/go/issues/25922
import (
	_ "github.com/client9/misspell/cmd/misspell"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "mvdan.cc/gofumpt/gofumports"
  _ "github.com/securego/gosec/v2/cmd/gosec"
	_ "golang.org/x/tools/cmd/goimports"
	_ "github.com/sqs/goreturns"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/go-critic/go-critic/cmd/gocritic"
)
