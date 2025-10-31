package ies

import "rrc/utils"

// FeatureSetUplink-v1610-intraFreqDAPS-UL-r16-dummy1 ::= ENUMERATED
type FeaturesetuplinkV1610IntrafreqdapsUlR16Dummy1 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610IntrafreqdapsUlR16Dummy1EnumeratedNothing = iota
	FeaturesetuplinkV1610IntrafreqdapsUlR16Dummy1EnumeratedSupported
)
