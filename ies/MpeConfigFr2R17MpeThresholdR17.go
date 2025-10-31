package ies

import "rrc/utils"

// MPE-Config-FR2-r17-mpe-Threshold-r17 ::= ENUMERATED
type MpeConfigFr2R17MpeThresholdR17 struct {
	Value utils.ENUMERATED
}

const (
	MpeConfigFr2R17MpeThresholdR17EnumeratedNothing = iota
	MpeConfigFr2R17MpeThresholdR17EnumeratedDb3
	MpeConfigFr2R17MpeThresholdR17EnumeratedDb6
	MpeConfigFr2R17MpeThresholdR17EnumeratedDb9
	MpeConfigFr2R17MpeThresholdR17EnumeratedDb12
)
