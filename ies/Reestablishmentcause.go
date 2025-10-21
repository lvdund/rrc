package ies

import "rrc/utils"

// ReestablishmentCause ::= ENUMERATED
type Reestablishmentcause struct {
	Value utils.ENUMERATED
}

const (
	ReestablishmentcauseReconfigurationfailure = 0
	ReestablishmentcauseHandoverfailure        = 1
	ReestablishmentcauseOtherfailure           = 2
	ReestablishmentcauseSpare1                 = 3
)
