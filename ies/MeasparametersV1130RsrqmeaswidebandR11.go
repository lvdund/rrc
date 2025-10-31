package ies

import "rrc/utils"

// MeasParameters-v1130-rsrqMeasWideband-r11 ::= ENUMERATED
type MeasparametersV1130RsrqmeaswidebandR11 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1130RsrqmeaswidebandR11EnumeratedNothing = iota
	MeasparametersV1130RsrqmeaswidebandR11EnumeratedSupported
)
