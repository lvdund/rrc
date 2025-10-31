package ies

import "rrc/utils"

// AdditionalPCIIndex-r17 ::= utils.INTEGER (1..maxNrofAdditionalPCI-r17)
type AdditionalpciindexR17 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofAdditionalPCIR17`
}
