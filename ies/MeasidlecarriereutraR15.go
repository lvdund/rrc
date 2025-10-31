package ies

// MeasIdleCarrierEUTRA-r15 ::= SEQUENCE
// Extensible
type MeasidlecarriereutraR15 struct {
	CarrierfreqR15          ArfcnValueeutraR9
	AllowedmeasbandwidthR15 Allowedmeasbandwidth
	ValidityareaR15         *CelllistR15
	MeascelllistR15         *CelllistR15
	Reportquantities        MeasidlecarriereutraR15Reportquantities
	QualitythresholdR15     *MeasidlecarriereutraR15QualitythresholdR15
}
