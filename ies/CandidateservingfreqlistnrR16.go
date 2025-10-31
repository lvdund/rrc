package ies

// CandidateServingFreqListNR-r16 ::= SEQUENCE OF ARFCN-ValueNR
// SIZE (1..maxFreqIDC-r16)
type CandidateservingfreqlistnrR16 struct {
	Value []ArfcnValuenr `lb:1,ub:maxFreqIDCR16`
}
