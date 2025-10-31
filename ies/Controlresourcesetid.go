package ies

import "rrc/utils"

// ControlResourceSetId ::= utils.INTEGER (0..maxNrofControlResourceSets-1)
type Controlresourcesetid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofControlResourceSets1`
}
