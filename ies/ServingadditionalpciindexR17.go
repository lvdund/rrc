package ies

import "rrc/utils"

// ServingAdditionalPCIIndex-r17 ::= utils.INTEGER (0..maxNrofAdditionalPCI-r17)
type ServingadditionalpciindexR17 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofAdditionalPCIR17`
}
