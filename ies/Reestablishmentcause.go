package ies

import "rrc/utils"

// ReestablishmentCause ::= ENUMERATED
type Reestablishmentcause struct {
	Value utils.ENUMERATED
}

const (
	ReestablishmentcauseEnumeratedNothing = iota
	ReestablishmentcauseEnumeratedReconfigurationfailure
	ReestablishmentcauseEnumeratedHandoverfailure
	ReestablishmentcauseEnumeratedOtherfailure
	ReestablishmentcauseEnumeratedSpare1
)
