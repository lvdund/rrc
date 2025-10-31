package ies

import "rrc/utils"

// FeatureSetDownlink-v1700-timeDurationForQCL-v1710-scs-480kHz ::= ENUMERATED
type FeaturesetdownlinkV1700TimedurationforqclV1710Scs480khz struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1700TimedurationforqclV1710Scs480khzEnumeratedNothing = iota
	FeaturesetdownlinkV1700TimedurationforqclV1710Scs480khzEnumeratedS56
	FeaturesetdownlinkV1700TimedurationforqclV1710Scs480khzEnumeratedS112
)
