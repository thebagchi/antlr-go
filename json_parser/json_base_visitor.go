// Code generated from json.g4 by ANTLR 4.10.1. DO NOT EDIT.

package json_parser // json
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasejsonVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasejsonVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasejsonVisitor) VisitDict(ctx *DictContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasejsonVisitor) VisitPairs(ctx *PairsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasejsonVisitor) VisitPair(ctx *PairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasejsonVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasejsonVisitor) VisitItems(ctx *ItemsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasejsonVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasejsonVisitor) VisitBoolean(ctx *BooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasejsonVisitor) VisitNull(ctx *NullContext) interface{} {
	return v.VisitChildren(ctx)
}
