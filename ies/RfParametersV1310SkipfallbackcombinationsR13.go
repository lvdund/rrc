package ies

import "rrc/utils"

// RF-Parameters-v1310-skipFallbackCombinations-r13 ::= ENUMERATED
type RfParametersV1310SkipfallbackcombinationsR13 struct {
	Value utils.ENUMERATED
}

const (
	RfParametersV1310SkipfallbackcombinationsR13EnumeratedNothing = iota
	RfParametersV1310SkipfallbackcombinationsR13EnumeratedSupported
)
