package ies

import "rrc/utils"

// PUCCH-Resource-intraSlotFrequencyHopping ::= ENUMERATED
type PucchResourceIntraslotfrequencyhopping struct {
	Value utils.ENUMERATED
}

const (
	PucchResourceIntraslotfrequencyhoppingEnumeratedNothing = iota
	PucchResourceIntraslotfrequencyhoppingEnumeratedEnabled
)
