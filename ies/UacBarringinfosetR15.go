package ies

import "rrc/utils"

// UAC-BarringInfoSet-r15 ::= SEQUENCE
type UacBarringinfosetR15 struct {
	UacBarringfactorR15            UacBarringinfosetR15UacBarringfactorR15
	UacBarringtimeR15              UacBarringinfosetR15UacBarringtimeR15
	UacBarringforaccessidentityR15 utils.BITSTRING `lb:7,ub:7`
}
