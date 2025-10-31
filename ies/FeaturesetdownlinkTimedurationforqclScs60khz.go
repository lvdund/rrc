package ies

import "rrc/utils"

// FeatureSetDownlink-timeDurationForQCL-scs-60kHz ::= ENUMERATED
type FeaturesetdownlinkTimedurationforqclScs60khz struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkTimedurationforqclScs60khzEnumeratedNothing = iota
	FeaturesetdownlinkTimedurationforqclScs60khzEnumeratedS7
	FeaturesetdownlinkTimedurationforqclScs60khzEnumeratedS14
	FeaturesetdownlinkTimedurationforqclScs60khzEnumeratedS28
)
