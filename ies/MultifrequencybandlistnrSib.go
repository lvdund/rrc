package ies

// MultiFrequencyBandListNR-SIB ::= SEQUENCE OF NR-MultiBandInfo
// SIZE (1.. maxNrofMultiBands)
type MultifrequencybandlistnrSib struct {
	Value []NrMultibandinfo `lb:1,ub:maxNrofMultiBands`
}
