package ies

import "rrc/utils"

// BandClassPriorityListHRPD ::= SEQUENCE OF BandClassPriorityHRPD
// SIZE (1..maxCDMA-BandClass)
type Bandclassprioritylisthrpd struct {
	Value utils.Sequence[Bandclasspriorityhrpd]
}
