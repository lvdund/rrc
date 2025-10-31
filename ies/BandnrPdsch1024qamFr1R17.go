package ies

import "rrc/utils"

// BandNR-pdsch-1024QAM-FR1-r17 ::= ENUMERATED
type BandnrPdsch1024qamFr1R17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrPdsch1024qamFr1R17EnumeratedNothing = iota
	BandnrPdsch1024qamFr1R17EnumeratedSupported
)
