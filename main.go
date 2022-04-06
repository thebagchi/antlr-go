package main

import (
	"antlr-go/json_parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

const content = `
{ 
    "Hello" : "World",
    "World" : "Hello",
    "Integer" : 12345,
    "Boolean" : true,
    "Null" : null,
    "Array" : [1,"1", true, false, null]
}
`

func main() {
	var (
		is    = antlr.NewInputStream(content)
		lexer = json_parser.NewjsonLexer(is)
	)
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Println(fmt.Sprintf("%s (%q)", lexer.SymbolicNames[t.GetTokenType()], t.GetText()))
	}
}
