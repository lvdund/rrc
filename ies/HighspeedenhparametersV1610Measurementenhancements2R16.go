package ies

import "rrc/utils"

// HighSpeedEnhParameters-v1610-measurementEnhancements2-r16 ::= ENUMERATED
type HighspeedenhparametersV1610Measurementenhancements2R16 struct {
	Value utils.ENUMERATED
}

const (
	HighspeedenhparametersV1610Measurementenhancements2R16EnumeratedNothing = iota
	HighspeedenhparametersV1610Measurementenhancements2R16EnumeratedSupported
)
