package ies

// CandidateServingFreqListNR-r15 ::= SEQUENCE OF ARFCN-ValueNR-r15
// SIZE (1..maxFreqIDC-r11)
type CandidateservingfreqlistnrR15 struct {
	Value []ArfcnValuenrR15 `lb:1,ub:maxFreqIDCR11`
}
