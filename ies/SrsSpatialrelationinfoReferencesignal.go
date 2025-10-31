package ies

// SRS-SpatialRelationInfo-referenceSignal ::= CHOICE
const (
	SrsSpatialrelationinfoReferencesignalChoiceNothing = iota
	SrsSpatialrelationinfoReferencesignalChoiceSsbIndex
	SrsSpatialrelationinfoReferencesignalChoiceCsiRsIndex
	SrsSpatialrelationinfoReferencesignalChoiceSrs
)

type SrsSpatialrelationinfoReferencesignal struct {
	Choice     uint64
	SsbIndex   *SsbIndex
	CsiRsIndex *NzpCsiRsResourceid
	Srs        *SrsSpatialrelationinfoReferencesignalSrs
}
