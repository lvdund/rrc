package ies

import "rrc/utils"

// GWUS-NumGroupsList-NB-r16 ::= SEQUENCE OF GWUS-NumGroups-NB-r16
// SIZE (1..maxGWUS-Resources-NB-r16)
type GwusNumgroupslistNbR16 struct {
	Value utils.Sequence[GwusNumgroupsNbR16]
}
