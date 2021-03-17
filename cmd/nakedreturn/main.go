package main

import (
	"github.com/gostaticanalysis/nakedreturn"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(nakedreturn.Analyzer) }
