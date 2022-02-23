package golinters

import (
	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis"
	"github.com/praetorian-inc/gokart/analyzers"
)

func NewGokart() *goanalysis.Linter {
	return goanalysis.NewLinter(
		"gokart",
		"A static analysis linter for securing Go code",
		analyzers.Analyzers,
		nil,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}
