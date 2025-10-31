package ies

import "rrc/utils"

// PUCCH-FormatConfig-interslotFrequencyHopping ::= ENUMERATED
type PucchFormatconfigInterslotfrequencyhopping struct {
	Value utils.ENUMERATED
}

const (
	PucchFormatconfigInterslotfrequencyhoppingEnumeratedNothing = iota
	PucchFormatconfigInterslotfrequencyhoppingEnumeratedEnabled
)
