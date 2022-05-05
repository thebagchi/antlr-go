// Code generated from json.g4 by ANTLR 4.10.1. DO NOT EDIT.

package json_parser // json
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasejsonListener is a complete listener for a parse tree produced by jsonParser.
type BasejsonListener struct{}

var _ jsonListener = &BasejsonListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasejsonListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasejsonListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasejsonListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasejsonListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BasejsonListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BasejsonListener) ExitStart(ctx *StartContext) {}

// EnterDict is called when production dict is entered.
func (s *BasejsonListener) EnterDict(ctx *DictContext) {}

// ExitDict is called when production dict is exited.
func (s *BasejsonListener) ExitDict(ctx *DictContext) {}

// EnterPairs is called when production pairs is entered.
func (s *BasejsonListener) EnterPairs(ctx *PairsContext) {}

// ExitPairs is called when production pairs is exited.
func (s *BasejsonListener) ExitPairs(ctx *PairsContext) {}

// EnterPair is called when production pair is entered.
func (s *BasejsonListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BasejsonListener) ExitPair(ctx *PairContext) {}

// EnterList is called when production list is entered.
func (s *BasejsonListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BasejsonListener) ExitList(ctx *ListContext) {}

// EnterItems is called when production items is entered.
func (s *BasejsonListener) EnterItems(ctx *ItemsContext) {}

// ExitItems is called when production items is exited.
func (s *BasejsonListener) ExitItems(ctx *ItemsContext) {}

// EnterValue is called when production value is entered.
func (s *BasejsonListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BasejsonListener) ExitValue(ctx *ValueContext) {}

// EnterBoolean is called when production boolean is entered.
func (s *BasejsonListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production boolean is exited.
func (s *BasejsonListener) ExitBoolean(ctx *BooleanContext) {}

// EnterNull is called when production null is entered.
func (s *BasejsonListener) EnterNull(ctx *NullContext) {}

// ExitNull is called when production null is exited.
func (s *BasejsonListener) ExitNull(ctx *NullContext) {}
