package ies

// SpatialRelationInfo-PDC-r17-referenceSignal ::= CHOICE
// Extensible
const (
	SpatialrelationinfoPdcR17ReferencesignalChoiceNothing = iota
	SpatialrelationinfoPdcR17ReferencesignalChoiceSsbIndex
	SpatialrelationinfoPdcR17ReferencesignalChoiceCsiRsIndex
	SpatialrelationinfoPdcR17ReferencesignalChoiceDlPrsPdc
	SpatialrelationinfoPdcR17ReferencesignalChoiceSrs
)

type SpatialrelationinfoPdcR17Referencesignal struct {
	Choice     uint64
	SsbIndex   *SsbIndex
	CsiRsIndex *NzpCsiRsResourceid
	DlPrsPdc   *NrDlPrsResourceidR17
	Srs        *SpatialrelationinfoPdcR17ReferencesignalSrs
}
