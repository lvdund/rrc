package ies

// PUSCH-PathlossReferenceRS-referenceSignal ::= CHOICE
const (
	PuschPathlossreferencersReferencesignalChoiceNothing = iota
	PuschPathlossreferencersReferencesignalChoiceSsbIndex
	PuschPathlossreferencersReferencesignalChoiceCsiRsIndex
)

type PuschPathlossreferencersReferencesignal struct {
	Choice     uint64
	SsbIndex   *SsbIndex
	CsiRsIndex *NzpCsiRsResourceid
}
