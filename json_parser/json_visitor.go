// Code generated from /Users/sandeep/Workspace/antlr-go/json_parser/json.g4 by ANTLR 4.9.2. DO NOT EDIT.

package json_parser // json
import "github.com/antlr/antlr4/runtime/Go/antlr"
// A complete Visitor for a parse tree produced by jsonParser.
type jsonVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by jsonParser#start.
	VisitStart(ctx *StartContext) interface{}

}