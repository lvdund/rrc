package ies

// SRS-SpatialRelationInfoPos-r16 ::= CHOICE
const (
	SrsSpatialrelationinfoposR16ChoiceNothing = iota
	SrsSpatialrelationinfoposR16ChoiceServingrsR16
	SrsSpatialrelationinfoposR16ChoiceSsbNcellR16
	SrsSpatialrelationinfoposR16ChoiceDlPrsR16
)

type SrsSpatialrelationinfoposR16 struct {
	Choice       uint64
	ServingrsR16 *SrsSpatialrelationinfoposR16ServingrsR16
	SsbNcellR16  *SsbInfoncellR16
	DlPrsR16     *DlPrsInfoR16
}
