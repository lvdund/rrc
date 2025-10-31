package ies

import "rrc/utils"

// BandNR-pdsch-1024QAM-2MIMO-FR1-r17 ::= ENUMERATED
type BandnrPdsch1024qam2mimoFr1R17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrPdsch1024qam2mimoFr1R17EnumeratedNothing = iota
	BandnrPdsch1024qam2mimoFr1R17EnumeratedSupported
)
