package nakedreturn

import (
	"go/ast"
	"go/types"
	"strings"

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
		(*ast.FuncDecl)(nil),
		(*ast.FuncLit)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			checkBody(pass, n.Body, pass.TypesInfo.TypeOf(n.Name))
		case *ast.FuncLit:
			checkBody(pass, n.Body, pass.TypesInfo.TypeOf(n))
		}
	})

	return nil, nil
}

func checkBody(pass *analysis.Pass, body *ast.BlockStmt, funType types.Type) {
	sig, _ := funType.(*types.Signature)
	if sig == nil || sig.Results().Len() == 0 {
		return
	}

	for _, stmt := range body.List {
		ret, _ := stmt.(*ast.ReturnStmt)
		if ret == nil || len(ret.Results) != 0 {
			continue
		}

		rets := make([]string, sig.Results().Len())
		for i := range rets {
			retv := sig.Results().At(i)
			name := retv.Name()
			if name == "_" {
				rets[i] = zeroValue(retv.Type())
			} else {
				rets[i] = name
			}
		}

		fix := analysis.SuggestedFix{
			Message: "add explicit return values",
			TextEdits: []analysis.TextEdit{{
				Pos:     ret.Pos(),
				End:     ret.End(),
				NewText: []byte("return " + strings.Join(rets, ", ")),
			}},
		}

		pass.Report(analysis.Diagnostic{
			Pos:            ret.Pos(),
			End:            ret.End(),
			Message:        "should not use naked return",
			SuggestedFixes: []analysis.SuggestedFix{fix},
		})
	}
}

func zeroValue(typ types.Type) string {
	switch utyp := typ.Underlying().(type) {
	case *types.Basic:
		switch {
		case utyp.Info() & types.IsNumeric != 0:
			return "0"
		case utyp.Info() & types.IsString != 0:
			return `""`
		default:
			return "false"
		}
	case *types.Struct, *types.Array:
		switch typ := typ.(type) {
		case *types.Named:
			return typ.Obj().Name() + "{}"
		default:
			return typ.String() + "{}"
		}
	default:
		return "nil"
	}
}
