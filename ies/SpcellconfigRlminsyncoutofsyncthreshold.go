package ies

import "rrc/utils"

// SpCellConfig-rlmInSyncOutOfSyncThreshold ::= ENUMERATED
type SpcellconfigRlminsyncoutofsyncthreshold struct {
	Value utils.ENUMERATED
}

const (
	SpcellconfigRlminsyncoutofsyncthresholdEnumeratedNothing = iota
	SpcellconfigRlminsyncoutofsyncthresholdEnumeratedN1
)
