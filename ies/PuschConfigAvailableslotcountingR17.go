package ies

import "rrc/utils"

// PUSCH-Config-availableSlotCounting-r17 ::= ENUMERATED
type PuschConfigAvailableslotcountingR17 struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigAvailableslotcountingR17EnumeratedNothing = iota
	PuschConfigAvailableslotcountingR17EnumeratedEnabled
)
