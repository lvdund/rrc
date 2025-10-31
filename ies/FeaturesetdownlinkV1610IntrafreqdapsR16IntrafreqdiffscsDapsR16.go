package ies

import "rrc/utils"

// FeatureSetDownlink-v1610-intraFreqDAPS-r16-intraFreqDiffSCS-DAPS-r16 ::= ENUMERATED
type FeaturesetdownlinkV1610IntrafreqdapsR16IntrafreqdiffscsDapsR16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1610IntrafreqdapsR16IntrafreqdiffscsDapsR16EnumeratedNothing = iota
	FeaturesetdownlinkV1610IntrafreqdapsR16IntrafreqdiffscsDapsR16EnumeratedSupported
)
