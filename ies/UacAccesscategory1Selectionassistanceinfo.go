package ies

import "rrc/utils"

// UAC-AccessCategory1-SelectionAssistanceInfo ::= ENUMERATED
type UacAccesscategory1Selectionassistanceinfo struct {
	Value utils.ENUMERATED
}

const (
	UacAccesscategory1SelectionassistanceinfoEnumeratedNothing = iota
	UacAccesscategory1SelectionassistanceinfoEnumeratedA
	UacAccesscategory1SelectionassistanceinfoEnumeratedB
	UacAccesscategory1SelectionassistanceinfoEnumeratedC
)
