package ies

import "rrc/utils"

// HighSpeedConfigFR2-r17-highSpeedDeploymentTypeFR2-r17 ::= ENUMERATED
type Highspeedconfigfr2R17Highspeeddeploymenttypefr2R17 struct {
	Value utils.ENUMERATED
}

const (
	Highspeedconfigfr2R17Highspeeddeploymenttypefr2R17EnumeratedNothing = iota
	Highspeedconfigfr2R17Highspeeddeploymenttypefr2R17EnumeratedUnidirectional
	Highspeedconfigfr2R17Highspeeddeploymenttypefr2R17EnumeratedBidirectional
)
