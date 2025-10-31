package ies

import "rrc/utils"

// TDD-Config-v1450-specialSubframePatterns-v1450 ::= ENUMERATED
type TddConfigV1450SpecialsubframepatternsV1450 struct {
	Value utils.ENUMERATED
}

const (
	TddConfigV1450SpecialsubframepatternsV1450EnumeratedNothing = iota
	TddConfigV1450SpecialsubframepatternsV1450EnumeratedSsp10_Crs_Lessdwpts
)
