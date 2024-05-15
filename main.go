// main.go
package main

import (
	"fmt"
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
