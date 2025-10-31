package ies

import "rrc/utils"

// FeatureSetUplink-v1610-intraFreqDAPS-UL-r16-intraFreqTwoTAGs-DAPS-r16 ::= ENUMERATED
type FeaturesetuplinkV1610IntrafreqdapsUlR16IntrafreqtwotagsDapsR16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610IntrafreqdapsUlR16IntrafreqtwotagsDapsR16EnumeratedNothing = iota
	FeaturesetuplinkV1610IntrafreqdapsUlR16IntrafreqtwotagsDapsR16EnumeratedSupported
)
