package ies

import "rrc/utils"

// FeatureSetDownlink-timeDurationForQCL-scs-120kHz ::= ENUMERATED
type FeaturesetdownlinkTimedurationforqclScs120khz struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkTimedurationforqclScs120khzEnumeratedNothing = iota
	FeaturesetdownlinkTimedurationforqclScs120khzEnumeratedS14
	FeaturesetdownlinkTimedurationforqclScs120khzEnumeratedS28
)
