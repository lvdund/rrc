package ies

import "rrc/utils"

// SL-UE-SelectedConfigRP-r16-sl-MaxNumPerReserve-r16 ::= ENUMERATED
type SlUeSelectedconfigrpR16SlMaxnumperreserveR16 struct {
	Value utils.ENUMERATED
}

const (
	SlUeSelectedconfigrpR16SlMaxnumperreserveR16EnumeratedNothing = iota
	SlUeSelectedconfigrpR16SlMaxnumperreserveR16EnumeratedN2
	SlUeSelectedconfigrpR16SlMaxnumperreserveR16EnumeratedN3
)
