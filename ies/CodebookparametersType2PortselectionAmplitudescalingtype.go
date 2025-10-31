package ies

import "rrc/utils"

// CodebookParameters-type2-PortSelection-amplitudeScalingType ::= ENUMERATED
type CodebookparametersType2PortselectionAmplitudescalingtype struct {
	Value utils.ENUMERATED
}

const (
	CodebookparametersType2PortselectionAmplitudescalingtypeEnumeratedNothing = iota
	CodebookparametersType2PortselectionAmplitudescalingtypeEnumeratedWideband
	CodebookparametersType2PortselectionAmplitudescalingtypeEnumeratedWidebandandsubband
)
