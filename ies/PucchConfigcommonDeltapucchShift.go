package ies

import "rrc/utils"

// PUCCH-ConfigCommon-deltaPUCCH-Shift ::= ENUMERATED
type PucchConfigcommonDeltapucchShift struct {
	Value utils.ENUMERATED
}

const (
	PucchConfigcommonDeltapucchShiftEnumeratedNothing = iota
	PucchConfigcommonDeltapucchShiftEnumeratedDs1
	PucchConfigcommonDeltapucchShiftEnumeratedDs2
	PucchConfigcommonDeltapucchShiftEnumeratedDs3
)
