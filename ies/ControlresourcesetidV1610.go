package ies

import "rrc/utils"

// ControlResourceSetId-v1610 ::= utils.INTEGER (maxNrofControlResourceSets..maxNrofControlResourceSets-1-r16)
type ControlresourcesetidV1610 struct {
	Value utils.INTEGER `lb:maxNrofControlResourceSets,ub:maxNrofControlResourceSets1R16`
}
