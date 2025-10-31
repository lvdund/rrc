package ies

import "rrc/utils"

// UAC-BarringInfoSet ::= SEQUENCE
type UacBarringinfoset struct {
	UacBarringfactor            UacBarringinfosetUacBarringfactor
	UacBarringtime              UacBarringinfosetUacBarringtime
	UacBarringforaccessidentity utils.BITSTRING `lb:7,ub:7`
}
