package ies

import "rrc/utils"

// MPE-Config-FR2-r16-mpe-Threshold-r16 ::= ENUMERATED
type MpeConfigFr2R16MpeThresholdR16 struct {
	Value utils.ENUMERATED
}

const (
	MpeConfigFr2R16MpeThresholdR16EnumeratedNothing = iota
	MpeConfigFr2R16MpeThresholdR16EnumeratedDb3
	MpeConfigFr2R16MpeThresholdR16EnumeratedDb6
	MpeConfigFr2R16MpeThresholdR16EnumeratedDb9
	MpeConfigFr2R16MpeThresholdR16EnumeratedDb12
)
