package ies

import "rrc/utils"

// BandNR-pusch-256QAM ::= ENUMERATED
type BandnrPusch256qam struct {
	Value utils.ENUMERATED
}

const (
	BandnrPusch256qamEnumeratedNothing = iota
	BandnrPusch256qamEnumeratedSupported
)
