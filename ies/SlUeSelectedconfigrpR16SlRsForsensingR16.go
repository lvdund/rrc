package ies

import "rrc/utils"

// SL-UE-SelectedConfigRP-r16-sl-RS-ForSensing-r16 ::= ENUMERATED
type SlUeSelectedconfigrpR16SlRsForsensingR16 struct {
	Value utils.ENUMERATED
}

const (
	SlUeSelectedconfigrpR16SlRsForsensingR16EnumeratedNothing = iota
	SlUeSelectedconfigrpR16SlRsForsensingR16EnumeratedPscch
	SlUeSelectedconfigrpR16SlRsForsensingR16EnumeratedPssch
)
