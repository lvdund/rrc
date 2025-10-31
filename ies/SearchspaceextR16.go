package ies

import "rrc/utils"

// SearchSpaceExt-r16 ::= SEQUENCE
// Extensible
type SearchspaceextR16 struct {
	ControlresourcesetidR16   *ControlresourcesetidR16
	SearchspacetypeR16        *SearchspaceextR16SearchspacetypeR16
	SearchspacegroupidlistR16 *[]utils.INTEGER `lb:1,ub:2`
	FreqmonitorlocationsR16   *utils.BITSTRING `lb:5,ub:5`
}
