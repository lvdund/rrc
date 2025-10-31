package ies

import "rrc/utils"

// FeatureSetUplink-v1610-intraFreqDAPS-UL-r16-dummy3 ::= ENUMERATED
type FeaturesetuplinkV1610IntrafreqdapsUlR16Dummy3 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610IntrafreqdapsUlR16Dummy3EnumeratedNothing = iota
	FeaturesetuplinkV1610IntrafreqdapsUlR16Dummy3EnumeratedShort
	FeaturesetuplinkV1610IntrafreqdapsUlR16Dummy3EnumeratedLong
)
