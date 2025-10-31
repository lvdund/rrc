package ies

import "rrc/utils"

// TDD-Config-v1430-specialSubframePatterns-v1430 ::= ENUMERATED
type TddConfigV1430SpecialsubframepatternsV1430 struct {
	Value utils.ENUMERATED
}

const (
	TddConfigV1430SpecialsubframepatternsV1430EnumeratedNothing = iota
	TddConfigV1430SpecialsubframepatternsV1430EnumeratedSsp10
)
