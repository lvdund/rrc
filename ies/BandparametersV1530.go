package ies

import "rrc/utils"

// BandParameters-v1530 ::= SEQUENCE
type BandparametersV1530 struct {
	UeTxantennaselectionSrs1t4rR15       *utils.ENUMERATED
	UeTxantennaselectionSrs2t4r2pairsR15 *utils.ENUMERATED
	UeTxantennaselectionSrs2t4r3pairsR15 *utils.ENUMERATED
	Dl1024qamR15                         *utils.ENUMERATED
	QclTypecOperationR15                 *utils.ENUMERATED
	QclCriBasedcsiReportingR15           *utils.ENUMERATED
	SttiSptBandparametersR15             *SttiSptBandparametersR15
}
