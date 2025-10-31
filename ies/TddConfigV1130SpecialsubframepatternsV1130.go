package ies

import "rrc/utils"

// TDD-Config-v1130-specialSubframePatterns-v1130 ::= ENUMERATED
type TddConfigV1130SpecialsubframepatternsV1130 struct {
	Value utils.ENUMERATED
}

const (
	TddConfigV1130SpecialsubframepatternsV1130EnumeratedNothing = iota
	TddConfigV1130SpecialsubframepatternsV1130EnumeratedSsp7
	TddConfigV1130SpecialsubframepatternsV1130EnumeratedSsp9
)
