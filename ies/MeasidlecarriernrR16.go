package ies

// MeasIdleCarrierNR-r16 ::= SEQUENCE
// Extensible
type MeasidlecarriernrR16 struct {
	CarrierfreqR16          ArfcnValuenr
	SsbsubcarrierspacingR16 Subcarrierspacing
	Frequencybandlist       *Multifrequencybandlistnr
	MeascelllistnrR16       *CelllistnrR16
	ReportquantitiesR16     MeasidlecarriernrR16ReportquantitiesR16
	QualitythresholdR16     *MeasidlecarriernrR16QualitythresholdR16
	SsbMeasconfigR16        *MeasidlecarriernrR16SsbMeasconfigR16
	BeammeasconfigidleR16   *BeammeasconfigidleNrR16
}
