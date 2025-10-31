package ies

import "rrc/utils"

// SL-UE-SelectedConfigRP-r16-sl-SensingWindow-r16 ::= ENUMERATED
type SlUeSelectedconfigrpR16SlSensingwindowR16 struct {
	Value utils.ENUMERATED
}

const (
	SlUeSelectedconfigrpR16SlSensingwindowR16EnumeratedNothing = iota
	SlUeSelectedconfigrpR16SlSensingwindowR16EnumeratedMs100
	SlUeSelectedconfigrpR16SlSensingwindowR16EnumeratedMs1100
)
