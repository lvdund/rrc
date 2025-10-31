package ies

// MeasIdleCarrierNR-r16 ::= SEQUENCE
// Extensible
type MeasidlecarriernrR16 struct {
	CarrierfreqnrR16        ArfcnValuenrR15
	SubcarrierspacingssbR16 MeasidlecarriernrR16SubcarrierspacingssbR16
	Frequencybandlist       *MultifrequencybandlistnrR15
	MeascelllistnrR16       *CelllistnrR16
	ReportquantitiesnrR16   MeasidlecarriernrR16ReportquantitiesnrR16
	QualitythresholdnrR16   *MeasidlecarriernrR16QualitythresholdnrR16
	SsbMeasconfigR16        *MeasidlecarriernrR16SsbMeasconfigR16
	BeammeasconfigidleR16   *BeammeasconfigidlenrR16
}
