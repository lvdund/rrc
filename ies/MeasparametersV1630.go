package ies

import "rrc/utils"

// MeasParameters-v1630 ::= SEQUENCE
type MeasparametersV1630 struct {
	NrIdleinactivebeammeasfr1R16 *utils.ENUMERATED
	NrIdleinactivebeammeasfr2R16 *utils.ENUMERATED
	CeMeasrssDedicatedsamerbsR16 *utils.ENUMERATED
}
