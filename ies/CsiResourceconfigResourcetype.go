package ies

import "rrc/utils"

// CSI-ResourceConfig-resourceType ::= ENUMERATED
type CsiResourceconfigResourcetype struct {
	Value utils.ENUMERATED
}

const (
	CsiResourceconfigResourcetypeEnumeratedNothing = iota
	CsiResourceconfigResourcetypeEnumeratedAperiodic
	CsiResourceconfigResourcetypeEnumeratedSemipersistent
	CsiResourceconfigResourcetypeEnumeratedPeriodic
)
