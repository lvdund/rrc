package ies

import "rrc/utils"

// UAC-Barring-NB-r16 ::= SEQUENCE
type UacBarringNbR16 struct {
	UacBarringpercatlistR16        *UacBarringpercatlistNbR16
	UacAc1SelectassistinfoR16      *UacAc1SelectassistinfoR15
	UacBarringforaccessidentityR16 utils.BITSTRING
}
