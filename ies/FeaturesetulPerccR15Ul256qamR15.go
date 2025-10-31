package ies

import "rrc/utils"

// FeatureSetUL-PerCC-r15-ul-256QAM-r15 ::= ENUMERATED
type FeaturesetulPerccR15Ul256qamR15 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetulPerccR15Ul256qamR15EnumeratedNothing = iota
	FeaturesetulPerccR15Ul256qamR15EnumeratedSupported
)
