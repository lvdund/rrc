package ies

import "rrc/utils"

// BandNR-max-HARQ-ProcessNumber-r17 ::= ENUMERATED
type BandnrMaxHarqProcessnumberR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrMaxHarqProcessnumberR17EnumeratedNothing = iota
	BandnrMaxHarqProcessnumberR17EnumeratedU16d32
	BandnrMaxHarqProcessnumberR17EnumeratedU32d16
	BandnrMaxHarqProcessnumberR17EnumeratedU32d32
)
