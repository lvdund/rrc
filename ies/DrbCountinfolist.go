package ies

import "rrc/utils"

// DRB-CountInfoList ::= SEQUENCE OF DRB-CountInfo
// SIZE (0..maxDRB)
type DrbCountinfolist struct {
	Value utils.Sequence[DrbCountinfo]
}
