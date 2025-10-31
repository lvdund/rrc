package ies

import "rrc/utils"

// SL-ZoneConfig-r16-sl-ZoneLength-r16 ::= ENUMERATED
type SlZoneconfigR16SlZonelengthR16 struct {
	Value utils.ENUMERATED
}

const (
	SlZoneconfigR16SlZonelengthR16EnumeratedNothing = iota
	SlZoneconfigR16SlZonelengthR16EnumeratedM5
	SlZoneconfigR16SlZonelengthR16EnumeratedM10
	SlZoneconfigR16SlZonelengthR16EnumeratedM20
	SlZoneconfigR16SlZonelengthR16EnumeratedM30
	SlZoneconfigR16SlZonelengthR16EnumeratedM40
	SlZoneconfigR16SlZonelengthR16EnumeratedM50
	SlZoneconfigR16SlZonelengthR16EnumeratedSpare2
	SlZoneconfigR16SlZonelengthR16EnumeratedSpare1
)
