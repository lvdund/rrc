package ies

// BandParameters-v1530 ::= SEQUENCE
type BandparametersV1530 struct {
	UeTxantennaselectionSrs1t4rR15       *BandparametersV1530UeTxantennaselectionSrs1t4rR15
	UeTxantennaselectionSrs2t4r2pairsR15 *BandparametersV1530UeTxantennaselectionSrs2t4r2pairsR15
	UeTxantennaselectionSrs2t4r3pairsR15 *BandparametersV1530UeTxantennaselectionSrs2t4r3pairsR15
	Dl1024qamR15                         *BandparametersV1530Dl1024qamR15
	QclTypecOperationR15                 *BandparametersV1530QclTypecOperationR15
	QclCriBasedcsiReportingR15           *BandparametersV1530QclCriBasedcsiReportingR15
	SttiSptBandparametersR15             *SttiSptBandparametersR15
}
