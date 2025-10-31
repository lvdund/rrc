package ies

import "rrc/utils"

// HighSpeedConfigFR2-r17 ::= SEQUENCE
// Extensible
type Highspeedconfigfr2R17 struct {
	Highspeedmeasflagfr2R17             *Highspeedconfigfr2R17Highspeedmeasflagfr2R17
	Highspeeddeploymenttypefr2R17       *Highspeedconfigfr2R17Highspeeddeploymenttypefr2R17
	HighspeedlargeonestepulTimingfr2R17 *utils.BOOLEAN
}
