package ies

import "rrc/utils"

// CodebookParameters-type2-amplitudeSubsetRestriction ::= ENUMERATED
type CodebookparametersType2Amplitudesubsetrestriction struct {
	Value utils.ENUMERATED
}

const (
	CodebookparametersType2AmplitudesubsetrestrictionEnumeratedNothing = iota
	CodebookparametersType2AmplitudesubsetrestrictionEnumeratedSupported
)
