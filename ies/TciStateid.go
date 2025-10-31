package ies

import "rrc/utils"

// TCI-StateId ::= utils.INTEGER (0..maxNrofTCI-States-1)
type TciStateid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofTCIStates1`
}
