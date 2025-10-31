package ies

import "rrc/utils"

// RSTD-InterFreqInfo-r10 ::= SEQUENCE
// Extensible
type RstdInterfreqinfoR10 struct {
	CarrierfreqR10   ArfcnValueeutra
	MeasprsOffsetR10 utils.INTEGER `lb:0,ub:39`
}
