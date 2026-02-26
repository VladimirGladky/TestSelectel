package SelectelTest

import (
	"github.com/VladimirGladky/SelectelTest/utils"
	"go/ast"
	"go/token"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "loglinter",
	Doc:  "check logs for established rules: log messages must start with lowercase letter, be in English only, not contain emojis or special characters, and not contain sensitive data",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
				methodName := sel.Sel.Name

				if utils.IsLogMethod(methodName) && !isFmtPackage(sel.X) {
					checkLogMessageFirstLetter(pass, call)
					checkLogMessageEnglishLanguage(pass, call)
					checkLogMessageSpecialChars(pass, call)
					checkLogMessageSensitiveData(pass, call)
				}
			}

			return true
		})
	}
	return nil, nil
}

func isFmtPackage(expr ast.Expr) bool {
	if ident, ok := expr.(*ast.Ident); ok {
		return ident.Name == "fmt" || ident.Name == "errors"
	}
	return false
}

func checkLogMessageFirstLetter(pass *analysis.Pass, call *ast.CallExpr) {
	if len(call.Args) == 0 {
		return
	}

	firstArg := call.Args[0]

	lit, ok := firstArg.(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return
	}

	message := strings.Trim(lit.Value, `"`+"`")

	if message == "" {
		return
	}

	firstRune := []rune(message)[0]

	if unicode.IsUpper(firstRune) {
		pass.Reportf(lit.Pos(), "log message should start with lowercase letter: %q", message)
	}
}

func checkLogMessageEnglishLanguage(pass *analysis.Pass, call *ast.CallExpr) {
	if len(call.Args) == 0 {
		return
	}

	firstArg := call.Args[0]

	lit, ok := firstArg.(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return
	}

	message := strings.Trim(lit.Value, `"`+"`")

	if message == "" {
		return
	}

	for _, r := range message {
		if unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsDigit(r) || unicode.IsSymbol(r) {
			continue
		}

		if unicode.IsLetter(r) && !utils.IsLatinLetter(r) {
			pass.Reportf(lit.Pos(), "log message must be in English only: %q", message)
			return
		}
	}
}

func checkLogMessageSpecialChars(pass *analysis.Pass, call *ast.CallExpr) {
	if len(call.Args) == 0 {
		return
	}

	firstArg := call.Args[0]

	lit, ok := firstArg.(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return
	}

	message := strings.Trim(lit.Value, `"`+"`")

	if message == "" {
		return
	}

	for _, r := range message {
		if utils.IsEmoji(r) {
			pass.Reportf(lit.Pos(), "log message must not contain emojis or special characters: %q", message)
			return
		}

		if utils.IsForbiddenPunctuation(r) {
			pass.Reportf(lit.Pos(), "log message must not contain emojis or special characters: %q", message)
			return
		}
	}
}

func checkLogMessageSensitiveData(pass *analysis.Pass, call *ast.CallExpr) {
	if len(call.Args) == 0 {
		return
	}

	firstArg := call.Args[0]

	lit, ok := firstArg.(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return
	}

	message := strings.Trim(lit.Value, `"`+"`")

	if message == "" {
		return
	}

	if utils.ContainsSensitiveData(message) {
		pass.Reportf(lit.Pos(), "log message must not contain sensitive data like passwords, tokens, or api keys: %q", message)
	}
}
