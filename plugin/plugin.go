package loglinter

import (
	"github.com/VladimirGladky/SelectelTest"

	"golang.org/x/tools/go/analysis"
)

// New is required for golangci-lint module plugins
func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{SelectelTest.Analyzer}, nil
}
