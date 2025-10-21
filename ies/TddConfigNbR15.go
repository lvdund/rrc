package ies

import "rrc/utils"

// TDD-Config-NB-r15 ::= SEQUENCE
type TddConfigNbR15 struct {
	SubframeassignmentR15      utils.ENUMERATED
	SpecialsubframepatternsR15 utils.ENUMERATED
}
