package ies

// CandidateServingFreqListEUTRA ::= SEQUENCE OF ARFCN-ValueEUTRA
// SIZE (1.. maxFreqIDC-MRDC)
type Candidateservingfreqlisteutra struct {
	Value []ArfcnValueeutra `lb:1,ub:maxFreqIDCMrdc`
}
