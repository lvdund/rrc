package ies

import "rrc/utils"

// BandParameters-v1530-ue-TxAntennaSelection-SRS-2T4R-3Pairs-r15 ::= ENUMERATED
type BandparametersV1530UeTxantennaselectionSrs2t4r3pairsR15 struct {
	Value utils.ENUMERATED
}

const (
	BandparametersV1530UeTxantennaselectionSrs2t4r3pairsR15EnumeratedNothing = iota
	BandparametersV1530UeTxantennaselectionSrs2t4r3pairsR15EnumeratedSupported
)
