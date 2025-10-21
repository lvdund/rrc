package ies

import "rrc/utils"

// MBMS-ServiceList-r13 ::= SEQUENCE OF MBMS-ServiceInfo-r13
// SIZE (0..maxMBMS-ServiceListPerUE-r13)
type MbmsServicelistR13 struct {
	Value utils.Sequence[MbmsServiceinfoR13]
}
