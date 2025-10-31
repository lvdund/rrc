package ies

import "rrc/utils"

// MIB-intraFreqReselection ::= ENUMERATED
type MibIntrafreqreselection struct {
	Value utils.ENUMERATED
}

const (
	MibIntrafreqreselectionEnumeratedNothing = iota
	MibIntrafreqreselectionEnumeratedAllowed
	MibIntrafreqreselectionEnumeratedNotallowed
)
