package ies

import "rrc/utils"

// SL-ZoneConfig-r14-zoneLength-r14 ::= ENUMERATED
type SlZoneconfigR14ZonelengthR14 struct {
	Value utils.ENUMERATED
}

const (
	SlZoneconfigR14ZonelengthR14EnumeratedNothing = iota
	SlZoneconfigR14ZonelengthR14EnumeratedM5
	SlZoneconfigR14ZonelengthR14EnumeratedM10
	SlZoneconfigR14ZonelengthR14EnumeratedM20
	SlZoneconfigR14ZonelengthR14EnumeratedM50
	SlZoneconfigR14ZonelengthR14EnumeratedM100
	SlZoneconfigR14ZonelengthR14EnumeratedM200
	SlZoneconfigR14ZonelengthR14EnumeratedM500
	SlZoneconfigR14ZonelengthR14EnumeratedSpare1
)
