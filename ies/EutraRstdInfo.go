package ies

import "rrc/utils"

// EUTRA-RSTD-Info ::= SEQUENCE
// Extensible
type EutraRstdInfo struct {
	Carrierfreq   ArfcnValueeutra
	MeasprsOffset utils.INTEGER `lb:0,ub:39`
}
