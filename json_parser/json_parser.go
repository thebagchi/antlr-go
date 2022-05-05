// Code generated from json.g4 by ANTLR 4.10.1. DO NOT EDIT.

package json_parser // json
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type jsonParser struct {
	*antlr.BaseParser
}

var jsonParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jsonParserInit() {
	staticData := &jsonParserStaticData
	staticData.literalNames = []string{
		"", "'true'", "'false'", "'null'", "'{'", "'}'", "':'", "'['", "']'",
		"','",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "CURLY_START", "CURLY_END", "COLON", "SQUARE_START",
		"SQUARE_END", "COMMA", "STRING", "NUMBER", "WHITESPACE",
	}
	staticData.ruleNames = []string{
		"start", "dict", "pairs", "pair", "list", "items", "value", "boolean",
		"null",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 12, 102, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 31, 8, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 42, 8, 2, 10, 2,
		12, 2, 45, 9, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 60, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 5, 5, 71, 8, 5, 10, 5, 12, 5, 74, 9, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 3, 6, 92, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 98, 8, 7, 1,
		8, 1, 8, 1, 8, 0, 2, 4, 10, 9, 0, 2, 4, 6, 8, 10, 12, 14, 16, 0, 0, 102,
		0, 18, 1, 0, 0, 0, 2, 30, 1, 0, 0, 0, 4, 32, 1, 0, 0, 0, 6, 46, 1, 0, 0,
		0, 8, 59, 1, 0, 0, 0, 10, 61, 1, 0, 0, 0, 12, 91, 1, 0, 0, 0, 14, 97, 1,
		0, 0, 0, 16, 99, 1, 0, 0, 0, 18, 19, 3, 12, 6, 0, 19, 20, 5, 0, 0, 1, 20,
		21, 6, 0, -1, 0, 21, 1, 1, 0, 0, 0, 22, 23, 5, 4, 0, 0, 23, 24, 3, 4, 2,
		0, 24, 25, 5, 5, 0, 0, 25, 26, 6, 1, -1, 0, 26, 31, 1, 0, 0, 0, 27, 28,
		5, 4, 0, 0, 28, 29, 5, 4, 0, 0, 29, 31, 6, 1, -1, 0, 30, 22, 1, 0, 0, 0,
		30, 27, 1, 0, 0, 0, 31, 3, 1, 0, 0, 0, 32, 33, 6, 2, -1, 0, 33, 34, 3,
		6, 3, 0, 34, 35, 6, 2, -1, 0, 35, 43, 1, 0, 0, 0, 36, 37, 10, 2, 0, 0,
		37, 38, 5, 9, 0, 0, 38, 39, 3, 6, 3, 0, 39, 40, 6, 2, -1, 0, 40, 42, 1,
		0, 0, 0, 41, 36, 1, 0, 0, 0, 42, 45, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 43,
		44, 1, 0, 0, 0, 44, 5, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 46, 47, 5, 10, 0,
		0, 47, 48, 5, 6, 0, 0, 48, 49, 3, 12, 6, 0, 49, 50, 6, 3, -1, 0, 50, 7,
		1, 0, 0, 0, 51, 52, 5, 7, 0, 0, 52, 53, 3, 10, 5, 0, 53, 54, 5, 8, 0, 0,
		54, 55, 6, 4, -1, 0, 55, 60, 1, 0, 0, 0, 56, 57, 5, 7, 0, 0, 57, 58, 5,
		8, 0, 0, 58, 60, 6, 4, -1, 0, 59, 51, 1, 0, 0, 0, 59, 56, 1, 0, 0, 0, 60,
		9, 1, 0, 0, 0, 61, 62, 6, 5, -1, 0, 62, 63, 3, 12, 6, 0, 63, 64, 6, 5,
		-1, 0, 64, 72, 1, 0, 0, 0, 65, 66, 10, 2, 0, 0, 66, 67, 5, 9, 0, 0, 67,
		68, 3, 12, 6, 0, 68, 69, 6, 5, -1, 0, 69, 71, 1, 0, 0, 0, 70, 65, 1, 0,
		0, 0, 71, 74, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 11,
		1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 75, 76, 5, 10, 0, 0, 76, 92, 6, 6, -1,
		0, 77, 78, 5, 11, 0, 0, 78, 92, 6, 6, -1, 0, 79, 80, 3, 2, 1, 0, 80, 81,
		6, 6, -1, 0, 81, 92, 1, 0, 0, 0, 82, 83, 3, 8, 4, 0, 83, 84, 6, 6, -1,
		0, 84, 92, 1, 0, 0, 0, 85, 86, 3, 14, 7, 0, 86, 87, 6, 6, -1, 0, 87, 92,
		1, 0, 0, 0, 88, 89, 3, 16, 8, 0, 89, 90, 6, 6, -1, 0, 90, 92, 1, 0, 0,
		0, 91, 75, 1, 0, 0, 0, 91, 77, 1, 0, 0, 0, 91, 79, 1, 0, 0, 0, 91, 82,
		1, 0, 0, 0, 91, 85, 1, 0, 0, 0, 91, 88, 1, 0, 0, 0, 92, 13, 1, 0, 0, 0,
		93, 94, 5, 1, 0, 0, 94, 98, 6, 7, -1, 0, 95, 96, 5, 2, 0, 0, 96, 98, 6,
		7, -1, 0, 97, 93, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 98, 15, 1, 0, 0, 0, 99,
		100, 5, 3, 0, 0, 100, 17, 1, 0, 0, 0, 6, 30, 43, 59, 72, 91, 97,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// jsonParserInit initializes any static state used to implement jsonParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewjsonParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JsonParserInit() {
	staticData := &jsonParserStaticData
	staticData.once.Do(jsonParserInit)
}

// NewjsonParser produces a new parser instance for the optional input antlr.TokenStream.
func NewjsonParser(input antlr.TokenStream) *jsonParser {
	JsonParserInit()
	this := new(jsonParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &jsonParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "json.g4"

	return this
}

// jsonParser tokens.
const (
	jsonParserEOF          = antlr.TokenEOF
	jsonParserT__0         = 1
	jsonParserT__1         = 2
	jsonParserT__2         = 3
	jsonParserCURLY_START  = 4
	jsonParserCURLY_END    = 5
	jsonParserCOLON        = 6
	jsonParserSQUARE_START = 7
	jsonParserSQUARE_END   = 8
	jsonParserCOMMA        = 9
	jsonParserSTRING       = 10
	jsonParserNUMBER       = 11
	jsonParserWHITESPACE   = 12
)

// jsonParser rules.
const (
	jsonParserRULE_start   = 0
	jsonParserRULE_dict    = 1
	jsonParserRULE_pairs   = 2
	jsonParserRULE_pair    = 3
	jsonParserRULE_list    = 4
	jsonParserRULE_items   = 5
	jsonParserRULE_value   = 6
	jsonParserRULE_boolean = 7
	jsonParserRULE_null    = 8
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetP1 returns the p1 rule contexts.
	GetP1() IValueContext

	// SetP1 sets the p1 rule contexts.
	SetP1(IValueContext)

	// GetV returns the v attribute.
	GetV() interface{}

	// SetV sets the v attribute.
	SetV(interface{})

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      interface{}
	p1     IValueContext
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

func (s *StartContext) GetP1() IValueContext { return s.p1 }

func (s *StartContext) SetP1(v IValueContext) { s.p1 = v }

func (s *StartContext) GetV() interface{} { return s.v }

func (s *StartContext) SetV(v interface{}) { s.v = v }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(jsonParserEOF, 0)
}

func (s *StartContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
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
	this := p
	_ = this

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
		p.SetState(18)

		var _x = p.Value()

		localctx.(*StartContext).p1 = _x
	}
	{
		p.SetState(19)
		p.Match(jsonParserEOF)
	}

	localctx.(*StartContext).v = localctx.(*StartContext).GetP1().GetV()

	return localctx
}

// IDictContext is an interface to support dynamic dispatch.
type IDictContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetP1 returns the p1 rule contexts.
	GetP1() IPairsContext

	// SetP1 sets the p1 rule contexts.
	SetP1(IPairsContext)

	// GetV returns the v attribute.
	GetV() map[string]interface{}

	// SetV sets the v attribute.
	SetV(map[string]interface{})

	// IsDictContext differentiates from other interfaces.
	IsDictContext()
}

type DictContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      map[string]interface{}
	p1     IPairsContext
}

func NewEmptyDictContext() *DictContext {
	var p = new(DictContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jsonParserRULE_dict
	return p
}

func (*DictContext) IsDictContext() {}

func NewDictContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictContext {
	var p = new(DictContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jsonParserRULE_dict

	return p
}

func (s *DictContext) GetParser() antlr.Parser { return s.parser }

func (s *DictContext) GetP1() IPairsContext { return s.p1 }

func (s *DictContext) SetP1(v IPairsContext) { s.p1 = v }

func (s *DictContext) GetV() map[string]interface{} { return s.v }

func (s *DictContext) SetV(v map[string]interface{}) { s.v = v }

func (s *DictContext) AllCURLY_START() []antlr.TerminalNode {
	return s.GetTokens(jsonParserCURLY_START)
}

func (s *DictContext) CURLY_START(i int) antlr.TerminalNode {
	return s.GetToken(jsonParserCURLY_START, i)
}

func (s *DictContext) CURLY_END() antlr.TerminalNode {
	return s.GetToken(jsonParserCURLY_END, 0)
}

func (s *DictContext) Pairs() IPairsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPairsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPairsContext)
}

func (s *DictContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.EnterDict(s)
	}
}

func (s *DictContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.ExitDict(s)
	}
}

func (s *DictContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case jsonVisitor:
		return t.VisitDict(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *jsonParser) Dict() (localctx IDictContext) {
	this := p
	_ = this

	localctx = NewDictContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, jsonParserRULE_dict)

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

	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.Match(jsonParserCURLY_START)
		}
		{
			p.SetState(23)

			var _x = p.pairs(0)

			localctx.(*DictContext).p1 = _x
		}
		{
			p.SetState(24)
			p.Match(jsonParserCURLY_END)
		}

		localctx.(*DictContext).v = localctx.(*DictContext).GetP1().GetV()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(27)
			p.Match(jsonParserCURLY_START)
		}
		{
			p.SetState(28)
			p.Match(jsonParserCURLY_START)
		}

		localctx.(*DictContext).v = map[string]interface{}{
			// EMPTY
		}

	}

	return localctx
}

// IPairsContext is an interface to support dynamic dispatch.
type IPairsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetP1 returns the p1 rule contexts.
	GetP1() IPairsContext

	// GetP3 returns the p3 rule contexts.
	GetP3() IPairContext

	// GetP2 returns the p2 rule contexts.
	GetP2() IPairContext

	// SetP1 sets the p1 rule contexts.
	SetP1(IPairsContext)

	// SetP3 sets the p3 rule contexts.
	SetP3(IPairContext)

	// SetP2 sets the p2 rule contexts.
	SetP2(IPairContext)

	// GetV returns the v attribute.
	GetV() map[string]interface{}

	// SetV sets the v attribute.
	SetV(map[string]interface{})

	// IsPairsContext differentiates from other interfaces.
	IsPairsContext()
}

type PairsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      map[string]interface{}
	p1     IPairsContext
	p3     IPairContext
	p2     IPairContext
}

func NewEmptyPairsContext() *PairsContext {
	var p = new(PairsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jsonParserRULE_pairs
	return p
}

func (*PairsContext) IsPairsContext() {}

func NewPairsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairsContext {
	var p = new(PairsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jsonParserRULE_pairs

	return p
}

func (s *PairsContext) GetParser() antlr.Parser { return s.parser }

func (s *PairsContext) GetP1() IPairsContext { return s.p1 }

func (s *PairsContext) GetP3() IPairContext { return s.p3 }

func (s *PairsContext) GetP2() IPairContext { return s.p2 }

func (s *PairsContext) SetP1(v IPairsContext) { s.p1 = v }

func (s *PairsContext) SetP3(v IPairContext) { s.p3 = v }

func (s *PairsContext) SetP2(v IPairContext) { s.p2 = v }

func (s *PairsContext) GetV() map[string]interface{} { return s.v }

func (s *PairsContext) SetV(v map[string]interface{}) { s.v = v }

func (s *PairsContext) Pair() IPairContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPairContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *PairsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(jsonParserCOMMA, 0)
}

func (s *PairsContext) Pairs() IPairsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPairsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPairsContext)
}

func (s *PairsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.EnterPairs(s)
	}
}

func (s *PairsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.ExitPairs(s)
	}
}

func (s *PairsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case jsonVisitor:
		return t.VisitPairs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *jsonParser) Pairs() (localctx IPairsContext) {
	return p.pairs(0)
}

func (p *jsonParser) pairs(_p int) (localctx IPairsContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPairsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPairsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, jsonParserRULE_pairs, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)

		var _x = p.Pair()

		localctx.(*PairsContext).p3 = _x
	}

	localctx.(*PairsContext).v = localctx.(*PairsContext).GetP3().GetV()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPairsContext(p, _parentctx, _parentState)
			localctx.(*PairsContext).p1 = _prevctx
			p.PushNewRecursionContext(localctx, _startState, jsonParserRULE_pairs)
			p.SetState(36)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(37)
				p.Match(jsonParserCOMMA)
			}
			{
				p.SetState(38)

				var _x = p.Pair()

				localctx.(*PairsContext).p2 = _x
			}

			localctx.(*PairsContext).v = localctx.(*PairsContext).GetP1().GetV()
			for k, v := range localctx.(*PairsContext).GetP2().GetV() {
				localctx.(*PairsContext).v[k] = v
			}

		}
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetP1 returns the p1 token.
	GetP1() antlr.Token

	// SetP1 sets the p1 token.
	SetP1(antlr.Token)

	// GetP2 returns the p2 rule contexts.
	GetP2() IValueContext

	// SetP2 sets the p2 rule contexts.
	SetP2(IValueContext)

	// GetV returns the v attribute.
	GetV() map[string]interface{}

	// SetV sets the v attribute.
	SetV(map[string]interface{})

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      map[string]interface{}
	p1     antlr.Token
	p2     IValueContext
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jsonParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jsonParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) GetP1() antlr.Token { return s.p1 }

func (s *PairContext) SetP1(v antlr.Token) { s.p1 = v }

func (s *PairContext) GetP2() IValueContext { return s.p2 }

func (s *PairContext) SetP2(v IValueContext) { s.p2 = v }

func (s *PairContext) GetV() map[string]interface{} { return s.v }

func (s *PairContext) SetV(v map[string]interface{}) { s.v = v }

func (s *PairContext) COLON() antlr.TerminalNode {
	return s.GetToken(jsonParserCOLON, 0)
}

func (s *PairContext) STRING() antlr.TerminalNode {
	return s.GetToken(jsonParserSTRING, 0)
}

func (s *PairContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.ExitPair(s)
	}
}

func (s *PairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case jsonVisitor:
		return t.VisitPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *jsonParser) Pair() (localctx IPairContext) {
	this := p
	_ = this

	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, jsonParserRULE_pair)

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
		p.SetState(46)

		var _m = p.Match(jsonParserSTRING)

		localctx.(*PairContext).p1 = _m
	}
	{
		p.SetState(47)
		p.Match(jsonParserCOLON)
	}
	{
		p.SetState(48)

		var _x = p.Value()

		localctx.(*PairContext).p2 = _x
	}

	fmt.Println(localctx.(*PairContext).GetP1())
	fmt.Println(localctx.(*PairContext).GetP2().GetV())
	localctx.(*PairContext).v = map[string]interface{}{
		(func() string {
			if localctx.(*PairContext).GetP1() == nil {
				return ""
			} else {
				return localctx.(*PairContext).GetP1().GetText()
			}
		}()): localctx.(*PairContext).GetP2().GetV(),
	}

	return localctx
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetP1 returns the p1 rule contexts.
	GetP1() IItemsContext

	// SetP1 sets the p1 rule contexts.
	SetP1(IItemsContext)

	// GetV returns the v attribute.
	GetV() []interface{}

	// SetV sets the v attribute.
	SetV([]interface{})

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      []interface{}
	p1     IItemsContext
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jsonParserRULE_list
	return p
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jsonParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) GetP1() IItemsContext { return s.p1 }

func (s *ListContext) SetP1(v IItemsContext) { s.p1 = v }

func (s *ListContext) GetV() []interface{} { return s.v }

func (s *ListContext) SetV(v []interface{}) { s.v = v }

func (s *ListContext) SQUARE_START() antlr.TerminalNode {
	return s.GetToken(jsonParserSQUARE_START, 0)
}

func (s *ListContext) SQUARE_END() antlr.TerminalNode {
	return s.GetToken(jsonParserSQUARE_END, 0)
}

func (s *ListContext) Items() IItemsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItemsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItemsContext)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.ExitList(s)
	}
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case jsonVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *jsonParser) List() (localctx IListContext) {
	this := p
	_ = this

	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, jsonParserRULE_list)

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

	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(51)
			p.Match(jsonParserSQUARE_START)
		}
		{
			p.SetState(52)

			var _x = p.items(0)

			localctx.(*ListContext).p1 = _x
		}
		{
			p.SetState(53)
			p.Match(jsonParserSQUARE_END)
		}

		localctx.(*ListContext).v = localctx.(*ListContext).GetP1().GetV()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.Match(jsonParserSQUARE_START)
		}
		{
			p.SetState(57)
			p.Match(jsonParserSQUARE_END)
		}

		localctx.(*ListContext).v = []interface{}{
			// EMPTY
		}

	}

	return localctx
}

// IItemsContext is an interface to support dynamic dispatch.
type IItemsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetP1 returns the p1 rule contexts.
	GetP1() IItemsContext

	// GetP3 returns the p3 rule contexts.
	GetP3() IValueContext

	// GetP2 returns the p2 rule contexts.
	GetP2() IValueContext

	// SetP1 sets the p1 rule contexts.
	SetP1(IItemsContext)

	// SetP3 sets the p3 rule contexts.
	SetP3(IValueContext)

	// SetP2 sets the p2 rule contexts.
	SetP2(IValueContext)

	// GetV returns the v attribute.
	GetV() []interface{}

	// SetV sets the v attribute.
	SetV([]interface{})

	// IsItemsContext differentiates from other interfaces.
	IsItemsContext()
}

type ItemsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      []interface{}
	p1     IItemsContext
	p3     IValueContext
	p2     IValueContext
}

func NewEmptyItemsContext() *ItemsContext {
	var p = new(ItemsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jsonParserRULE_items
	return p
}

func (*ItemsContext) IsItemsContext() {}

func NewItemsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ItemsContext {
	var p = new(ItemsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jsonParserRULE_items

	return p
}

func (s *ItemsContext) GetParser() antlr.Parser { return s.parser }

func (s *ItemsContext) GetP1() IItemsContext { return s.p1 }

func (s *ItemsContext) GetP3() IValueContext { return s.p3 }

func (s *ItemsContext) GetP2() IValueContext { return s.p2 }

func (s *ItemsContext) SetP1(v IItemsContext) { s.p1 = v }

func (s *ItemsContext) SetP3(v IValueContext) { s.p3 = v }

func (s *ItemsContext) SetP2(v IValueContext) { s.p2 = v }

func (s *ItemsContext) GetV() []interface{} { return s.v }

func (s *ItemsContext) SetV(v []interface{}) { s.v = v }

func (s *ItemsContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ItemsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(jsonParserCOMMA, 0)
}

func (s *ItemsContext) Items() IItemsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItemsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItemsContext)
}

func (s *ItemsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItemsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ItemsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.EnterItems(s)
	}
}

func (s *ItemsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.ExitItems(s)
	}
}

func (s *ItemsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case jsonVisitor:
		return t.VisitItems(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *jsonParser) Items() (localctx IItemsContext) {
	return p.items(0)
}

func (p *jsonParser) items(_p int) (localctx IItemsContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewItemsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IItemsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, jsonParserRULE_items, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)

		var _x = p.Value()

		localctx.(*ItemsContext).p3 = _x
	}

	localctx.(*ItemsContext).v = []interface{}{
		localctx.(*ItemsContext).GetP3().GetV(),
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewItemsContext(p, _parentctx, _parentState)
			localctx.(*ItemsContext).p1 = _prevctx
			p.PushNewRecursionContext(localctx, _startState, jsonParserRULE_items)
			p.SetState(65)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(66)
				p.Match(jsonParserCOMMA)
			}
			{
				p.SetState(67)

				var _x = p.Value()

				localctx.(*ItemsContext).p2 = _x
			}

			localctx.(*ItemsContext).v = localctx.(*ItemsContext).GetP1().GetV()
			localctx.(*ItemsContext).v = append(localctx.(*ItemsContext).v, localctx.(*ItemsContext).GetP2().GetV())

		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetP1 returns the p1 token.
	GetP1() antlr.Token

	// GetP2 returns the p2 token.
	GetP2() antlr.Token

	// SetP1 sets the p1 token.
	SetP1(antlr.Token)

	// SetP2 sets the p2 token.
	SetP2(antlr.Token)

	// GetP3 returns the p3 rule contexts.
	GetP3() IDictContext

	// GetP4 returns the p4 rule contexts.
	GetP4() IListContext

	// GetP5 returns the p5 rule contexts.
	GetP5() IBooleanContext

	// SetP3 sets the p3 rule contexts.
	SetP3(IDictContext)

	// SetP4 sets the p4 rule contexts.
	SetP4(IListContext)

	// SetP5 sets the p5 rule contexts.
	SetP5(IBooleanContext)

	// GetV returns the v attribute.
	GetV() interface{}

	// SetV sets the v attribute.
	SetV(interface{})

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      interface{}
	p1     antlr.Token
	p2     antlr.Token
	p3     IDictContext
	p4     IListContext
	p5     IBooleanContext
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jsonParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jsonParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) GetP1() antlr.Token { return s.p1 }

func (s *ValueContext) GetP2() antlr.Token { return s.p2 }

func (s *ValueContext) SetP1(v antlr.Token) { s.p1 = v }

func (s *ValueContext) SetP2(v antlr.Token) { s.p2 = v }

func (s *ValueContext) GetP3() IDictContext { return s.p3 }

func (s *ValueContext) GetP4() IListContext { return s.p4 }

func (s *ValueContext) GetP5() IBooleanContext { return s.p5 }

func (s *ValueContext) SetP3(v IDictContext) { s.p3 = v }

func (s *ValueContext) SetP4(v IListContext) { s.p4 = v }

func (s *ValueContext) SetP5(v IBooleanContext) { s.p5 = v }

func (s *ValueContext) GetV() interface{} { return s.v }

func (s *ValueContext) SetV(v interface{}) { s.v = v }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(jsonParserSTRING, 0)
}

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(jsonParserNUMBER, 0)
}

func (s *ValueContext) Dict() IDictContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDictContext)
}

func (s *ValueContext) List() IListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *ValueContext) Boolean() IBooleanContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanContext)
}

func (s *ValueContext) Null() INullContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INullContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INullContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case jsonVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *jsonParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, jsonParserRULE_value)

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

	p.SetState(91)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case jsonParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)

			var _m = p.Match(jsonParserSTRING)

			localctx.(*ValueContext).p1 = _m
		}

		localctx.(*ValueContext).v = (func() string {
			if localctx.(*ValueContext).GetP1() == nil {
				return ""
			} else {
				return localctx.(*ValueContext).GetP1().GetText()
			}
		}())

	case jsonParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)

			var _m = p.Match(jsonParserNUMBER)

			localctx.(*ValueContext).p2 = _m
		}

		v, _ := strconv.ParseFloat((func() string {
			if localctx.(*ValueContext).GetP2() == nil {
				return ""
			} else {
				return localctx.(*ValueContext).GetP2().GetText()
			}
		}()), 64)
		localctx.(*ValueContext).v = v

	case jsonParserCURLY_START:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(79)

			var _x = p.Dict()

			localctx.(*ValueContext).p3 = _x
		}

		localctx.(*ValueContext).v = localctx.(*ValueContext).GetP3().GetV()

	case jsonParserSQUARE_START:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(82)

			var _x = p.List()

			localctx.(*ValueContext).p4 = _x
		}

		localctx.(*ValueContext).v = localctx.(*ValueContext).GetP4().GetV()

	case jsonParserT__0, jsonParserT__1:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(85)

			var _x = p.Boolean()

			localctx.(*ValueContext).p5 = _x
		}

		localctx.(*ValueContext).v = localctx.(*ValueContext).GetP5().GetV()

	case jsonParserT__2:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(88)
			p.Null()
		}

		localctx.(*ValueContext).v = nil

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBooleanContext is an interface to support dynamic dispatch.
type IBooleanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v attribute.
	GetV() bool

	// SetV sets the v attribute.
	SetV(bool)

	// IsBooleanContext differentiates from other interfaces.
	IsBooleanContext()
}

type BooleanContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      bool
}

func NewEmptyBooleanContext() *BooleanContext {
	var p = new(BooleanContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jsonParserRULE_boolean
	return p
}

func (*BooleanContext) IsBooleanContext() {}

func NewBooleanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanContext {
	var p = new(BooleanContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jsonParserRULE_boolean

	return p
}

func (s *BooleanContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanContext) GetV() bool { return s.v }

func (s *BooleanContext) SetV(v bool) { s.v = v }

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.EnterBoolean(s)
	}
}

func (s *BooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.ExitBoolean(s)
	}
}

func (s *BooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case jsonVisitor:
		return t.VisitBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *jsonParser) Boolean() (localctx IBooleanContext) {
	this := p
	_ = this

	localctx = NewBooleanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, jsonParserRULE_boolean)

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

	p.SetState(97)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case jsonParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)
			p.Match(jsonParserT__0)
		}

		localctx.(*BooleanContext).v = true

	case jsonParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(95)
			p.Match(jsonParserT__1)
		}

		localctx.(*BooleanContext).v = false

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INullContext is an interface to support dynamic dispatch.
type INullContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNullContext differentiates from other interfaces.
	IsNullContext()
}

type NullContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNullContext() *NullContext {
	var p = new(NullContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jsonParserRULE_null
	return p
}

func (*NullContext) IsNullContext() {}

func NewNullContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NullContext {
	var p = new(NullContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jsonParserRULE_null

	return p
}

func (s *NullContext) GetParser() antlr.Parser { return s.parser }
func (s *NullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.EnterNull(s)
	}
}

func (s *NullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.ExitNull(s)
	}
}

func (s *NullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case jsonVisitor:
		return t.VisitNull(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *jsonParser) Null() (localctx INullContext) {
	this := p
	_ = this

	localctx = NewNullContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, jsonParserRULE_null)

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
		p.SetState(99)
		p.Match(jsonParserT__2)
	}

	return localctx
}

func (p *jsonParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *PairsContext = nil
		if localctx != nil {
			t = localctx.(*PairsContext)
		}
		return p.Pairs_Sempred(t, predIndex)

	case 5:
		var t *ItemsContext = nil
		if localctx != nil {
			t = localctx.(*ItemsContext)
		}
		return p.Items_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *jsonParser) Pairs_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *jsonParser) Items_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
