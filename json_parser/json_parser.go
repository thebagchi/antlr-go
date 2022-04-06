// Code generated from /Users/sandeep/Workspace/antlr-go/json_parser/json.g4 by ANTLR 4.9.2. DO NOT EDIT.

package json_parser // json
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 22, 8, 4, 
	2, 9, 2, 3, 2, 3, 2, 3, 2, 3, 2, 2, 2, 3, 2, 2, 2, 2, 6, 2, 4, 3, 2, 2, 
	2, 4, 5, 7, 3, 2, 2, 5, 6, 7, 2, 2, 3, 6, 3, 3, 2, 2, 2, 2,
}
var literalNames = []string{
	"", "", "", "", "", "", "", "", "", "'true'", "'false'", "'null'", "'{'", 
	"'}'", "'['", "']'", "':'", "','",
}
var symbolicNames = []string{
	"", "Value", "Dict", "List", "String", "Number", "Boolean", "Null", "Pair", 
	"TRUE_SYM", "FALSE_SYM", "NULL_SYM", "CURLY_START", "CURLY_END", "SQUARE_START", 
	"SQUARE_END", "COLON", "COMMA", "NUMBER", "STRING", "WHITESPACE",
}

var ruleNames = []string{
	"start",
}
type jsonParser struct {
	*antlr.BaseParser
}

// NewjsonParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *jsonParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewjsonParser(input antlr.TokenStream) *jsonParser {
	this := new(jsonParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "json.g4"

	return this
}


// jsonParser tokens.
const (
	jsonParserEOF = antlr.TokenEOF
	jsonParserValue = 1
	jsonParserDict = 2
	jsonParserList = 3
	jsonParserString = 4
	jsonParserNumber = 5
	jsonParserBoolean = 6
	jsonParserNull = 7
	jsonParserPair = 8
	jsonParserTRUE_SYM = 9
	jsonParserFALSE_SYM = 10
	jsonParserNULL_SYM = 11
	jsonParserCURLY_START = 12
	jsonParserCURLY_END = 13
	jsonParserSQUARE_START = 14
	jsonParserSQUARE_END = 15
	jsonParserCOLON = 16
	jsonParserCOMMA = 17
	jsonParserNUMBER = 18
	jsonParserSTRING = 19
	jsonParserWHITESPACE = 20
)

// jsonParserRULE_start is the jsonParser rule.
const jsonParserRULE_start = 0


// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jsonParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jsonParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Value() antlr.TerminalNode {
	return s.GetToken(jsonParserValue, 0)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(jsonParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.ExitStart(s)
	}
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case jsonVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *jsonParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, jsonParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(2)
		p.Match(jsonParserValue)
	}
	{
		p.SetState(3)
		p.Match(jsonParserEOF)
	}



	return localctx
}


