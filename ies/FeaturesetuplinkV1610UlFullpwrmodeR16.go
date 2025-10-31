package ies

import "rrc/utils"

// FeatureSetUplink-v1610-ul-FullPwrMode-r16 ::= ENUMERATED
type FeaturesetuplinkV1610UlFullpwrmodeR16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610UlFullpwrmodeR16EnumeratedNothing = iota
	FeaturesetuplinkV1610UlFullpwrmodeR16EnumeratedSupported
)
