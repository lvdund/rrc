package ies

import "rrc/utils"

// DRB-InfoListSCG-r12 ::= SEQUENCE OF DRB-InfoSCG-r12
// SIZE (1..maxDRB)
type DrbInfolistscgR12 struct {
	Value utils.Sequence[DrbInfoscgR12]
}
