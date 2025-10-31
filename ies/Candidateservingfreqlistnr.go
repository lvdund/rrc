package ies

// CandidateServingFreqListNR ::= SEQUENCE OF ARFCN-ValueNR
// SIZE (1.. maxFreqIDC-MRDC)
type Candidateservingfreqlistnr struct {
	Value []ArfcnValuenr `lb:1,ub:maxFreqIDCMrdc`
}
