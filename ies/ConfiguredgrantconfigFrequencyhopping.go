package ies

import "rrc/utils"

// ConfiguredGrantConfig-frequencyHopping ::= ENUMERATED
type ConfiguredgrantconfigFrequencyhopping struct {
	Value utils.ENUMERATED
}

const (
	ConfiguredgrantconfigFrequencyhoppingEnumeratedNothing = iota
	ConfiguredgrantconfigFrequencyhoppingEnumeratedIntraslot
	ConfiguredgrantconfigFrequencyhoppingEnumeratedInterslot
)
