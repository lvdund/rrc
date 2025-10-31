package ies

// MeasIdleCarrierEUTRA-r16 ::= SEQUENCE
// Extensible
type MeasidlecarriereutraR16 struct {
	CarrierfreqeutraR16      ArfcnValueeutra
	AllowedmeasbandwidthR16  EutraAllowedmeasbandwidth
	MeascelllisteutraR16     *CelllisteutraR16
	ReportquantitieseutraR16 MeasidlecarriereutraR16ReportquantitieseutraR16
	QualitythresholdeutraR16 *MeasidlecarriereutraR16QualitythresholdeutraR16
}
