package ies

import "rrc/utils"

// SRB-ToAddModList ::= SEQUENCE OF SRB-ToAddMod
// SIZE (1..2)
type SrbToaddmodlist struct {
	Value utils.Sequence[SrbToaddmod]
}
