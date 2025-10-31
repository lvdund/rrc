package ies

import "rrc/utils"

// ReestablishmentCause-NB-r13 ::= ENUMERATED
type ReestablishmentcauseNbR13 struct {
	Value utils.ENUMERATED
}

const (
	ReestablishmentcauseNbR13EnumeratedNothing = iota
	ReestablishmentcauseNbR13EnumeratedReconfigurationfailure
	ReestablishmentcauseNbR13EnumeratedOtherfailure
	ReestablishmentcauseNbR13EnumeratedSpare2
	ReestablishmentcauseNbR13EnumeratedSpare1
)
