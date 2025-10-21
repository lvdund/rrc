package ies

import "rrc/utils"

// DRB-ToAddModList ::= SEQUENCE OF DRB-ToAddMod
// SIZE (1..maxDRB)
type DrbToaddmodlist struct {
	Value utils.Sequence[DrbToaddmod]
}
