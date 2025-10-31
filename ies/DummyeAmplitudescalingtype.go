package ies

import "rrc/utils"

// DummyE-amplitudeScalingType ::= ENUMERATED
type DummyeAmplitudescalingtype struct {
	Value utils.ENUMERATED
}

const (
	DummyeAmplitudescalingtypeEnumeratedNothing = iota
	DummyeAmplitudescalingtypeEnumeratedWideband
	DummyeAmplitudescalingtypeEnumeratedWidebandandsubband
)
