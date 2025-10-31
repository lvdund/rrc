package ies

// EUTRA-RSTD-InfoList ::= SEQUENCE OF EUTRA-RSTD-Info
// SIZE (1..maxInterRAT-RSTD-Freq)
type EutraRstdInfolist struct {
	Value []EutraRstdInfo `lb:1,ub:maxInterRATRstdFreq`
}
