package main

import (
	"github.com/VladimirGladky/SelectelTest"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(SelectelTest.Analyzer)
}
