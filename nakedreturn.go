package nakedreturn

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "nakedreturn finds naked returns"

// Analyzer finds naked returns.
var Analyzer = &analysis.Analyzer{
	Name: "nakedreturn",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.ReturnStmt)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.ReturnStmt:
			if len(n.Results) == 0 {
				pass.Reportf(n.Pos(), "should not use naked return")
			}
		}
	})

	return nil, nil
}
