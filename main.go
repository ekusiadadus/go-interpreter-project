// main.go
package main

import (
	"fmt"

	"github.com/ekusiadadus/go-interpreter-project/evaluate"
	"github.com/ekusiadadus/go-interpreter-project/lexer"
	"github.com/ekusiadadus/go-interpreter-project/parser"
)

func main() {
	source := `
        // ソースコードを書く
    `
	tokens := lexer.Lex(source)
	ast := parser.Parse(tokens)
	result := evaluate.Evaluate(ast)
	fmt.Println(result)
}
