package ies

import "rrc/utils"

// BandParameters-v1530-ue-TxAntennaSelection-SRS-1T4R-r15 ::= ENUMERATED
type BandparametersV1530UeTxantennaselectionSrs1t4rR15 struct {
	Value utils.ENUMERATED
}

const (
	BandparametersV1530UeTxantennaselectionSrs1t4rR15EnumeratedNothing = iota
	BandparametersV1530UeTxantennaselectionSrs1t4rR15EnumeratedSupported
)
