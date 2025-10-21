package ies

import "rrc/utils"

// MeasObjectToAddModList ::= SEQUENCE OF MeasObjectToAddMod
// SIZE (1..maxObjectId)
type Measobjecttoaddmodlist struct {
	Value utils.Sequence[Measobjecttoaddmod]
}
