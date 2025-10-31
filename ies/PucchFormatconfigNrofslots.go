package ies

import "rrc/utils"

// PUCCH-FormatConfig-nrofSlots ::= ENUMERATED
type PucchFormatconfigNrofslots struct {
	Value utils.ENUMERATED
}

const (
	PucchFormatconfigNrofslotsEnumeratedNothing = iota
	PucchFormatconfigNrofslotsEnumeratedN2
	PucchFormatconfigNrofslotsEnumeratedN4
	PucchFormatconfigNrofslotsEnumeratedN8
)
