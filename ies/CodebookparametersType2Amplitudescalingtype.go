package ies

import "rrc/utils"

// CodebookParameters-type2-amplitudeScalingType ::= ENUMERATED
type CodebookparametersType2Amplitudescalingtype struct {
	Value utils.ENUMERATED
}

const (
	CodebookparametersType2AmplitudescalingtypeEnumeratedNothing = iota
	CodebookparametersType2AmplitudescalingtypeEnumeratedWideband
	CodebookparametersType2AmplitudescalingtypeEnumeratedWidebandandsubband
)
