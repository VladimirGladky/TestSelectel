package main

import (
	"SelectelTest"

	"golang.org/x/tools/go/analysis"
)

// AnalyzerPlugin is required for golangci-lint module plugins
type AnalyzerPlugin struct{}

func (*AnalyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		SelectelTest.Analyzer,
	}
}

// New is required for golangci-lint module plugins
func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{SelectelTest.Analyzer}, nil
}
