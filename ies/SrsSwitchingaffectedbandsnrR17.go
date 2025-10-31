package ies

import "rrc/utils"

// SRS-SwitchingAffectedBandsNR-r17 ::= BIT STRING (SIZE (1..maxSimultaneousBands))
type SrsSwitchingaffectedbandsnrR17 struct {
	Value utils.BITSTRING `lb:1,ub:maxSimultaneousBands`
}
