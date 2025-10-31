package ies

// PUCCH-SpatialRelationInfo-referenceSignal ::= CHOICE
const (
	PucchSpatialrelationinfoReferencesignalChoiceNothing = iota
	PucchSpatialrelationinfoReferencesignalChoiceSsbIndex
	PucchSpatialrelationinfoReferencesignalChoiceCsiRsIndex
	PucchSpatialrelationinfoReferencesignalChoiceSrs
)

type PucchSpatialrelationinfoReferencesignal struct {
	Choice     uint64
	SsbIndex   *SsbIndex
	CsiRsIndex *NzpCsiRsResourceid
	Srs        *PucchSrs
}
