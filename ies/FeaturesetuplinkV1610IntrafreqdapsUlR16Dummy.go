package ies

import "rrc/utils"

// FeatureSetUplink-v1610-intraFreqDAPS-UL-r16-dummy ::= ENUMERATED
type FeaturesetuplinkV1610IntrafreqdapsUlR16Dummy struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610IntrafreqdapsUlR16DummyEnumeratedNothing = iota
	FeaturesetuplinkV1610IntrafreqdapsUlR16DummyEnumeratedSupported
)
