package ies

// SL-MinT2ValueList-r15 ::= SEQUENCE OF SL-MinT2Value-r15
// SIZE (1..maxSL-Prio-r13)
type SlMint2valuelistR15 struct {
	Value []SlMint2valueR15 `lb:1,ub:maxSLPrioR13`
}
