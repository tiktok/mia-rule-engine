// Code generated from parser/cmpl.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type cmplLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var CmplLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cmpllexerLexerInit() {
	staticData := &CmplLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'DEFAULT '", "';'", "' WITH '", "' WHEN '", "' PRIORITY '", "':'",
		"','", "'.'", "'AND'", "'OR'", "'GRANT $scope AS '", "'EXEC $action AS '",
		"'ADD $fact AS '", "", "", "", "", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "AND", "OR", "SCOPE", "ACTION",
		"FACT", "ID", "STR", "NUM", "OPS", "LPAREN", "RPAREN", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "AND",
		"OR", "SCOPE", "ACTION", "FACT", "ID", "STR", "NUM", "OPS", "LPAREN",
		"RPAREN", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 20, 188, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 5, 13, 140, 8, 13, 10, 13,
		12, 13, 143, 9, 13, 1, 13, 1, 13, 5, 13, 147, 8, 13, 10, 13, 12, 13, 150,
		9, 13, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 156, 8, 14, 10, 14, 12, 14, 159,
		9, 14, 1, 14, 1, 14, 1, 15, 4, 15, 164, 8, 15, 11, 15, 12, 15, 165, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 176, 8, 16,
		1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 4, 19, 183, 8, 19, 11, 19, 12, 19, 184,
		1, 19, 1, 19, 0, 0, 20, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 1, 0, 7, 1, 0, 36, 36, 3, 0, 65, 90, 95, 95, 97,
		122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 34, 34, 92, 92, 1, 0,
		48, 57, 2, 0, 60, 60, 62, 62, 3, 0, 9, 10, 13, 13, 32, 32, 197, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1,
		0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17,
		1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0,
		25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0,
		0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0,
		0, 1, 41, 1, 0, 0, 0, 3, 50, 1, 0, 0, 0, 5, 52, 1, 0, 0, 0, 7, 59, 1, 0,
		0, 0, 9, 66, 1, 0, 0, 0, 11, 77, 1, 0, 0, 0, 13, 79, 1, 0, 0, 0, 15, 81,
		1, 0, 0, 0, 17, 83, 1, 0, 0, 0, 19, 87, 1, 0, 0, 0, 21, 90, 1, 0, 0, 0,
		23, 107, 1, 0, 0, 0, 25, 124, 1, 0, 0, 0, 27, 141, 1, 0, 0, 0, 29, 151,
		1, 0, 0, 0, 31, 163, 1, 0, 0, 0, 33, 175, 1, 0, 0, 0, 35, 177, 1, 0, 0,
		0, 37, 179, 1, 0, 0, 0, 39, 182, 1, 0, 0, 0, 41, 42, 5, 68, 0, 0, 42, 43,
		5, 69, 0, 0, 43, 44, 5, 70, 0, 0, 44, 45, 5, 65, 0, 0, 45, 46, 5, 85, 0,
		0, 46, 47, 5, 76, 0, 0, 47, 48, 5, 84, 0, 0, 48, 49, 5, 32, 0, 0, 49, 2,
		1, 0, 0, 0, 50, 51, 5, 59, 0, 0, 51, 4, 1, 0, 0, 0, 52, 53, 5, 32, 0, 0,
		53, 54, 5, 87, 0, 0, 54, 55, 5, 73, 0, 0, 55, 56, 5, 84, 0, 0, 56, 57,
		5, 72, 0, 0, 57, 58, 5, 32, 0, 0, 58, 6, 1, 0, 0, 0, 59, 60, 5, 32, 0,
		0, 60, 61, 5, 87, 0, 0, 61, 62, 5, 72, 0, 0, 62, 63, 5, 69, 0, 0, 63, 64,
		5, 78, 0, 0, 64, 65, 5, 32, 0, 0, 65, 8, 1, 0, 0, 0, 66, 67, 5, 32, 0,
		0, 67, 68, 5, 80, 0, 0, 68, 69, 5, 82, 0, 0, 69, 70, 5, 73, 0, 0, 70, 71,
		5, 79, 0, 0, 71, 72, 5, 82, 0, 0, 72, 73, 5, 73, 0, 0, 73, 74, 5, 84, 0,
		0, 74, 75, 5, 89, 0, 0, 75, 76, 5, 32, 0, 0, 76, 10, 1, 0, 0, 0, 77, 78,
		5, 58, 0, 0, 78, 12, 1, 0, 0, 0, 79, 80, 5, 44, 0, 0, 80, 14, 1, 0, 0,
		0, 81, 82, 5, 46, 0, 0, 82, 16, 1, 0, 0, 0, 83, 84, 5, 65, 0, 0, 84, 85,
		5, 78, 0, 0, 85, 86, 5, 68, 0, 0, 86, 18, 1, 0, 0, 0, 87, 88, 5, 79, 0,
		0, 88, 89, 5, 82, 0, 0, 89, 20, 1, 0, 0, 0, 90, 91, 5, 71, 0, 0, 91, 92,
		5, 82, 0, 0, 92, 93, 5, 65, 0, 0, 93, 94, 5, 78, 0, 0, 94, 95, 5, 84, 0,
		0, 95, 96, 5, 32, 0, 0, 96, 97, 5, 36, 0, 0, 97, 98, 5, 115, 0, 0, 98,
		99, 5, 99, 0, 0, 99, 100, 5, 111, 0, 0, 100, 101, 5, 112, 0, 0, 101, 102,
		5, 101, 0, 0, 102, 103, 5, 32, 0, 0, 103, 104, 5, 65, 0, 0, 104, 105, 5,
		83, 0, 0, 105, 106, 5, 32, 0, 0, 106, 22, 1, 0, 0, 0, 107, 108, 5, 69,
		0, 0, 108, 109, 5, 88, 0, 0, 109, 110, 5, 69, 0, 0, 110, 111, 5, 67, 0,
		0, 111, 112, 5, 32, 0, 0, 112, 113, 5, 36, 0, 0, 113, 114, 5, 97, 0, 0,
		114, 115, 5, 99, 0, 0, 115, 116, 5, 116, 0, 0, 116, 117, 5, 105, 0, 0,
		117, 118, 5, 111, 0, 0, 118, 119, 5, 110, 0, 0, 119, 120, 5, 32, 0, 0,
		120, 121, 5, 65, 0, 0, 121, 122, 5, 83, 0, 0, 122, 123, 5, 32, 0, 0, 123,
		24, 1, 0, 0, 0, 124, 125, 5, 65, 0, 0, 125, 126, 5, 68, 0, 0, 126, 127,
		5, 68, 0, 0, 127, 128, 5, 32, 0, 0, 128, 129, 5, 36, 0, 0, 129, 130, 5,
		102, 0, 0, 130, 131, 5, 97, 0, 0, 131, 132, 5, 99, 0, 0, 132, 133, 5, 116,
		0, 0, 133, 134, 5, 32, 0, 0, 134, 135, 5, 65, 0, 0, 135, 136, 5, 83, 0,
		0, 136, 137, 5, 32, 0, 0, 137, 26, 1, 0, 0, 0, 138, 140, 7, 0, 0, 0, 139,
		138, 1, 0, 0, 0, 140, 143, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 142,
		1, 0, 0, 0, 142, 144, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 144, 148, 7, 1,
		0, 0, 145, 147, 7, 2, 0, 0, 146, 145, 1, 0, 0, 0, 147, 150, 1, 0, 0, 0,
		148, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 28, 1, 0, 0, 0, 150, 148,
		1, 0, 0, 0, 151, 157, 5, 34, 0, 0, 152, 156, 8, 3, 0, 0, 153, 154, 5, 92,
		0, 0, 154, 156, 9, 0, 0, 0, 155, 152, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0,
		156, 159, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158,
		160, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160, 161, 5, 34, 0, 0, 161, 30,
		1, 0, 0, 0, 162, 164, 7, 4, 0, 0, 163, 162, 1, 0, 0, 0, 164, 165, 1, 0,
		0, 0, 165, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 32, 1, 0, 0, 0,
		167, 176, 5, 61, 0, 0, 168, 169, 5, 33, 0, 0, 169, 176, 5, 61, 0, 0, 170,
		176, 7, 5, 0, 0, 171, 172, 5, 60, 0, 0, 172, 176, 5, 61, 0, 0, 173, 174,
		5, 62, 0, 0, 174, 176, 5, 61, 0, 0, 175, 167, 1, 0, 0, 0, 175, 168, 1,
		0, 0, 0, 175, 170, 1, 0, 0, 0, 175, 171, 1, 0, 0, 0, 175, 173, 1, 0, 0,
		0, 176, 34, 1, 0, 0, 0, 177, 178, 5, 40, 0, 0, 178, 36, 1, 0, 0, 0, 179,
		180, 5, 41, 0, 0, 180, 38, 1, 0, 0, 0, 181, 183, 7, 6, 0, 0, 182, 181,
		1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 184, 185, 1, 0,
		0, 0, 185, 186, 1, 0, 0, 0, 186, 187, 6, 19, 0, 0, 187, 40, 1, 0, 0, 0,
		8, 0, 141, 148, 155, 157, 165, 175, 184, 1, 6, 0, 0,
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

// cmplLexerInit initializes any static state used to implement cmplLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewcmplLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CmplLexerInit() {
	staticData := &CmplLexerLexerStaticData
	staticData.once.Do(cmpllexerLexerInit)
}

// NewcmplLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewcmplLexer(input antlr.CharStream) *cmplLexer {
	CmplLexerInit()
	l := new(cmplLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &CmplLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "cmpl.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// cmplLexer tokens.
const (
	cmplLexerT__0   = 1
	cmplLexerT__1   = 2
	cmplLexerT__2   = 3
	cmplLexerT__3   = 4
	cmplLexerT__4   = 5
	cmplLexerT__5   = 6
	cmplLexerT__6   = 7
	cmplLexerT__7   = 8
	cmplLexerAND    = 9
	cmplLexerOR     = 10
	cmplLexerSCOPE  = 11
	cmplLexerACTION = 12
	cmplLexerFACT   = 13
	cmplLexerID     = 14
	cmplLexerSTR    = 15
	cmplLexerNUM    = 16
	cmplLexerOPS    = 17
	cmplLexerLPAREN = 18
	cmplLexerRPAREN = 19
	cmplLexerWS     = 20
)
