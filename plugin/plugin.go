package linters

import (
	"github.com/VladimirGladky/SelectelTest"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("loglinter", New)
}

type LogLinterPlugin struct{}

func New(settings any) (register.LinterPlugin, error) {
	return &LogLinterPlugin{}, nil
}

func (p *LogLinterPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{SelectelTest.Analyzer}, nil
}

func (p *LogLinterPlugin) GetLoadMode() string {
	return register.LoadModeSyntax
}
