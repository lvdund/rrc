package ies

// SRS-SpatialRelationInfoPos-r16-servingRS-r16-referenceSignal-r16 ::= CHOICE
const (
	SrsSpatialrelationinfoposR16ServingrsR16ReferencesignalR16ChoiceNothing = iota
	SrsSpatialrelationinfoposR16ServingrsR16ReferencesignalR16ChoiceSsbIndexservingR16
	SrsSpatialrelationinfoposR16ServingrsR16ReferencesignalR16ChoiceCsiRsIndexservingR16
	SrsSpatialrelationinfoposR16ServingrsR16ReferencesignalR16ChoiceSrsSpatialrelationR16
)

type SrsSpatialrelationinfoposR16ServingrsR16ReferencesignalR16 struct {
	Choice                uint64
	SsbIndexservingR16    *SsbIndex
	CsiRsIndexservingR16  *NzpCsiRsResourceid
	SrsSpatialrelationR16 *SrsSpatialrelationinfoposR16ServingrsR16ReferencesignalR16SrsSpatialrelationR16
}
