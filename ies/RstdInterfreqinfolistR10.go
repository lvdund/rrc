package ies

// RSTD-InterFreqInfoList-r10 ::= SEQUENCE OF RSTD-InterFreqInfo-r10
// SIZE (1..maxRSTD-Freq-r10)
type RstdInterfreqinfolistR10 struct {
	Value []RstdInterfreqinfoR10 `lb:1,ub:maxRSTDFreqR10`
}
