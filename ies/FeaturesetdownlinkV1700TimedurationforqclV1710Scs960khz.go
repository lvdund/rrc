package ies

import "rrc/utils"

// FeatureSetDownlink-v1700-timeDurationForQCL-v1710-scs-960kHz ::= ENUMERATED
type FeaturesetdownlinkV1700TimedurationforqclV1710Scs960khz struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1700TimedurationforqclV1710Scs960khzEnumeratedNothing = iota
	FeaturesetdownlinkV1700TimedurationforqclV1710Scs960khzEnumeratedS112
	FeaturesetdownlinkV1700TimedurationforqclV1710Scs960khzEnumeratedS224
)
