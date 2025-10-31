package ies

import "rrc/utils"

// BandNR-pdsch-256QAM-FR2 ::= ENUMERATED
type BandnrPdsch256qamFr2 struct {
	Value utils.ENUMERATED
}

const (
	BandnrPdsch256qamFr2EnumeratedNothing = iota
	BandnrPdsch256qamFr2EnumeratedSupported
)
