package ies

import "rrc/utils"

// SL-UE-SelectedConfigRP-r16-sl-MultiReserveResource-r16 ::= ENUMERATED
type SlUeSelectedconfigrpR16SlMultireserveresourceR16 struct {
	Value utils.ENUMERATED
}

const (
	SlUeSelectedconfigrpR16SlMultireserveresourceR16EnumeratedNothing = iota
	SlUeSelectedconfigrpR16SlMultireserveresourceR16EnumeratedEnabled
)
