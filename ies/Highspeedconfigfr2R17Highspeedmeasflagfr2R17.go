package ies

import "rrc/utils"

// HighSpeedConfigFR2-r17-highSpeedMeasFlagFR2-r17 ::= ENUMERATED
type Highspeedconfigfr2R17Highspeedmeasflagfr2R17 struct {
	Value utils.ENUMERATED
}

const (
	Highspeedconfigfr2R17Highspeedmeasflagfr2R17EnumeratedNothing = iota
	Highspeedconfigfr2R17Highspeedmeasflagfr2R17EnumeratedSet1
	Highspeedconfigfr2R17Highspeedmeasflagfr2R17EnumeratedSet2
)
