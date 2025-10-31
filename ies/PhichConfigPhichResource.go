package ies

import "rrc/utils"

// PHICH-Config-phich-Resource ::= ENUMERATED
type PhichConfigPhichResource struct {
	Value utils.ENUMERATED
}

const (
	PhichConfigPhichResourceEnumeratedNothing = iota
	PhichConfigPhichResourceEnumeratedOnesixth
	PhichConfigPhichResourceEnumeratedHalf
	PhichConfigPhichResourceEnumeratedOne
	PhichConfigPhichResourceEnumeratedTwo
)
