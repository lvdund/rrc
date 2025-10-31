package ies

import "rrc/utils"

// PHICH-Config-phich-Duration ::= ENUMERATED
type PhichConfigPhichDuration struct {
	Value utils.ENUMERATED
}

const (
	PhichConfigPhichDurationEnumeratedNothing = iota
	PhichConfigPhichDurationEnumeratedNormal
	PhichConfigPhichDurationEnumeratedExtended
)
