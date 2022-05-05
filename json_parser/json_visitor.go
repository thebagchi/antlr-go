// Code generated from json.g4 by ANTLR 4.10.1. DO NOT EDIT.

package json_parser // json
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by jsonParser.
type jsonVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by jsonParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by jsonParser#dict.
	VisitDict(ctx *DictContext) interface{}

	// Visit a parse tree produced by jsonParser#pairs.
	VisitPairs(ctx *PairsContext) interface{}

	// Visit a parse tree produced by jsonParser#pair.
	VisitPair(ctx *PairContext) interface{}

	// Visit a parse tree produced by jsonParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by jsonParser#items.
	VisitItems(ctx *ItemsContext) interface{}

	// Visit a parse tree produced by jsonParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by jsonParser#boolean.
	VisitBoolean(ctx *BooleanContext) interface{}

	// Visit a parse tree produced by jsonParser#null.
	VisitNull(ctx *NullContext) interface{}
}
