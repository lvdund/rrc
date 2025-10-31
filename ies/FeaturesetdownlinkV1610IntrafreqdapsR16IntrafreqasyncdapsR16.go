package ies

import "rrc/utils"

// FeatureSetDownlink-v1610-intraFreqDAPS-r16-intraFreqAsyncDAPS-r16 ::= ENUMERATED
type FeaturesetdownlinkV1610IntrafreqdapsR16IntrafreqasyncdapsR16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1610IntrafreqdapsR16IntrafreqasyncdapsR16EnumeratedNothing = iota
	FeaturesetdownlinkV1610IntrafreqdapsR16IntrafreqasyncdapsR16EnumeratedSupported
)
