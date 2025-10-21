package ies

import "rrc/utils"

// TDD-Config ::= SEQUENCE
type TddConfig struct {
	Subframeassignment      utils.ENUMERATED
	Specialsubframepatterns utils.ENUMERATED
}
