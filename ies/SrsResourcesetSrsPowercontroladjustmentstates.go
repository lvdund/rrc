package ies

import "rrc/utils"

// SRS-ResourceSet-srs-PowerControlAdjustmentStates ::= ENUMERATED
type SrsResourcesetSrsPowercontroladjustmentstates struct {
	Value utils.ENUMERATED
}

const (
	SrsResourcesetSrsPowercontroladjustmentstatesEnumeratedNothing = iota
	SrsResourcesetSrsPowercontroladjustmentstatesEnumeratedSameasfci2
	SrsResourcesetSrsPowercontroladjustmentstatesEnumeratedSeparateclosedloop
)
