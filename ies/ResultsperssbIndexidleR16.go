package ies

// ResultsPerSSB-IndexIdle-r16 ::= SEQUENCE
type ResultsperssbIndexidleR16 struct {
	SsbIndexR16   SsbIndex
	SsbResultsR16 *ResultsperssbIndexidleR16SsbResultsR16
}
