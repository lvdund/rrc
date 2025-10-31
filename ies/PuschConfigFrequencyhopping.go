package ies

import "rrc/utils"

// PUSCH-Config-frequencyHopping ::= ENUMERATED
type PuschConfigFrequencyhopping struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigFrequencyhoppingEnumeratedNothing = iota
	PuschConfigFrequencyhoppingEnumeratedIntraslot
	PuschConfigFrequencyhoppingEnumeratedInterslot
)
