package ies

// PUCCH-PathlossReferenceRS-referenceSignal ::= CHOICE
const (
	PucchPathlossreferencersReferencesignalChoiceNothing = iota
	PucchPathlossreferencersReferencesignalChoiceSsbIndex
	PucchPathlossreferencersReferencesignalChoiceCsiRsIndex
)

type PucchPathlossreferencersReferencesignal struct {
	Choice     uint64
	SsbIndex   *SsbIndex
	CsiRsIndex *NzpCsiRsResourceid
}
