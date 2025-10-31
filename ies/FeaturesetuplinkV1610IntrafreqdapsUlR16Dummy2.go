package ies

import "rrc/utils"

// FeatureSetUplink-v1610-intraFreqDAPS-UL-r16-dummy2 ::= ENUMERATED
type FeaturesetuplinkV1610IntrafreqdapsUlR16Dummy2 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610IntrafreqdapsUlR16Dummy2EnumeratedNothing = iota
	FeaturesetuplinkV1610IntrafreqdapsUlR16Dummy2EnumeratedSupported
)
