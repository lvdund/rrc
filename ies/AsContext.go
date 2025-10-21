package ies

import "rrc/utils"

// AS-Context ::= SEQUENCE
type AsContext struct {
	Reestablishmentinfo *Reestablishmentinfo
}
