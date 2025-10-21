package ies

import "rrc/utils"

// ReestablishmentCause-NB-r13 ::= ENUMERATED
type ReestablishmentcauseNbR13 struct {
	Value utils.ENUMERATED
}

const (
	ReestablishmentcauseNbR13Reconfigurationfailure = 0
	ReestablishmentcauseNbR13Otherfailure           = 1
	ReestablishmentcauseNbR13Spare2                 = 2
	ReestablishmentcauseNbR13Spare1                 = 3
)
