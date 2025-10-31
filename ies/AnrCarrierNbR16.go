package ies

import "rrc/utils"

// ANR-Carrier-NB-r16 ::= SEQUENCE
// Extensible
type AnrCarrierNbR16 struct {
	CarrierfreqindexR16 utils.INTEGER `lb:0,ub:maxFreq`
	BlackcelllistR16    *AnrBlackcelllistNbR16
}
