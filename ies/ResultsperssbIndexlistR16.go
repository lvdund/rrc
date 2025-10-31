package ies

// ResultsPerSSB-IndexList-r16 ::= SEQUENCE OF ResultsPerSSB-IndexIdle-r16
// SIZE (1..maxRS-IndexReport-r15)
type ResultsperssbIndexlistR16 struct {
	Value []ResultsperssbIndexidleR16 `lb:1,ub:maxRSIndexreportR15`
}
