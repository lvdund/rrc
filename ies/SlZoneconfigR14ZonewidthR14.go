package ies

import "rrc/utils"

// SL-ZoneConfig-r14-zoneWidth-r14 ::= ENUMERATED
type SlZoneconfigR14ZonewidthR14 struct {
	Value utils.ENUMERATED
}

const (
	SlZoneconfigR14ZonewidthR14EnumeratedNothing = iota
	SlZoneconfigR14ZonewidthR14EnumeratedM5
	SlZoneconfigR14ZonewidthR14EnumeratedM10
	SlZoneconfigR14ZonewidthR14EnumeratedM20
	SlZoneconfigR14ZonewidthR14EnumeratedM50
	SlZoneconfigR14ZonewidthR14EnumeratedM100
	SlZoneconfigR14ZonewidthR14EnumeratedM200
	SlZoneconfigR14ZonewidthR14EnumeratedM500
	SlZoneconfigR14ZonewidthR14EnumeratedSpare1
)
