// Code generated from json.g4 by ANTLR 4.10.1. DO NOT EDIT.

package json_parser // json
import "github.com/antlr/antlr4/runtime/Go/antlr"

// jsonListener is a complete listener for a parse tree produced by jsonParser.
type jsonListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterDict is called when entering the dict production.
	EnterDict(c *DictContext)

	// EnterPairs is called when entering the pairs production.
	EnterPairs(c *PairsContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterItems is called when entering the items production.
	EnterItems(c *ItemsContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterBoolean is called when entering the boolean production.
	EnterBoolean(c *BooleanContext)

	// EnterNull is called when entering the null production.
	EnterNull(c *NullContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitDict is called when exiting the dict production.
	ExitDict(c *DictContext)

	// ExitPairs is called when exiting the pairs production.
	ExitPairs(c *PairsContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitItems is called when exiting the items production.
	ExitItems(c *ItemsContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitBoolean is called when exiting the boolean production.
	ExitBoolean(c *BooleanContext)

	// ExitNull is called when exiting the null production.
	ExitNull(c *NullContext)
}
