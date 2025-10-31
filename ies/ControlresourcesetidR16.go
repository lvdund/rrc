package ies

import "rrc/utils"

// ControlResourceSetId-r16 ::= utils.INTEGER (0..maxNrofControlResourceSets-1-r16)
type ControlresourcesetidR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofControlResourceSets1R16`
}
