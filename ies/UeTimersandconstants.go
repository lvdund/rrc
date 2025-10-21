package ies

import "rrc/utils"

// UE-TimersAndConstants ::= SEQUENCE
// Extensible
type UeTimersandconstants struct {
	T300 utils.ENUMERATED
	T301 utils.ENUMERATED
	T310 utils.ENUMERATED
	N310 utils.ENUMERATED
	T311 utils.ENUMERATED
	N311 utils.ENUMERATED
}
