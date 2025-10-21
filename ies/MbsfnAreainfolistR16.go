package ies

import "rrc/utils"

// MBSFN-AreaInfoList-r16 ::= SEQUENCE OF MBSFN-AreaInfo-r16
// SIZE (1..maxMBSFN-Area)
type MbsfnAreainfolistR16 struct {
	Value utils.Sequence[MbsfnAreainfoR16]
}
