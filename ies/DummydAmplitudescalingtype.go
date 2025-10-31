package ies

import "rrc/utils"

// DummyD-amplitudeScalingType ::= ENUMERATED
type DummydAmplitudescalingtype struct {
	Value utils.ENUMERATED
}

const (
	DummydAmplitudescalingtypeEnumeratedNothing = iota
	DummydAmplitudescalingtypeEnumeratedWideband
	DummydAmplitudescalingtypeEnumeratedWidebandandsubband
)
