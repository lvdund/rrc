package ies

import "rrc/utils"

// RSTD-InterFreqInfoList-r10 ::= SEQUENCE OF RSTD-InterFreqInfo-r10
// SIZE (1..maxRSTD-Freq-r10)
type RstdInterfreqinfolistR10 struct {
	Value utils.Sequence[RstdInterfreqinfoR10]
}
