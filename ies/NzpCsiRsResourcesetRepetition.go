package ies

import "rrc/utils"

// NZP-CSI-RS-ResourceSet-repetition ::= ENUMERATED
type NzpCsiRsResourcesetRepetition struct {
	Value utils.ENUMERATED
}

const (
	NzpCsiRsResourcesetRepetitionEnumeratedNothing = iota
	NzpCsiRsResourcesetRepetitionEnumeratedOn
	NzpCsiRsResourcesetRepetitionEnumeratedOff
)
