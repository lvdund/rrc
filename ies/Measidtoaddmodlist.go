package ies

import "rrc/utils"

// MeasIdToAddModList ::= SEQUENCE OF MeasIdToAddMod
// SIZE (1..maxMeasId)
type Measidtoaddmodlist struct {
	Value []Measidtoaddmod
}
