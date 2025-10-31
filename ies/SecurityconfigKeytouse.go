package ies

import "rrc/utils"

// SecurityConfig-keyToUse ::= ENUMERATED
type SecurityconfigKeytouse struct {
	Value utils.ENUMERATED
}

const (
	SecurityconfigKeytouseEnumeratedNothing = iota
	SecurityconfigKeytouseEnumeratedMaster
	SecurityconfigKeytouseEnumeratedSecondary
)
