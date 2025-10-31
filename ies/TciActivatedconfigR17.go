package ies

import "rrc/utils"

// TCI-ActivatedConfig-r17 ::= SEQUENCE
type TciActivatedconfigR17 struct {
	PdcchTciR17 []TciStateid    `lb:1,ub:5`
	PdschTciR17 utils.BITSTRING `lb:1,ub:maxNrofTCIStates`
}
