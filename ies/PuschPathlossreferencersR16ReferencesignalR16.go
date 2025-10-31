package ies

// PUSCH-PathlossReferenceRS-r16-referenceSignal-r16 ::= CHOICE
const (
	PuschPathlossreferencersR16ReferencesignalR16ChoiceNothing = iota
	PuschPathlossreferencersR16ReferencesignalR16ChoiceSsbIndexR16
	PuschPathlossreferencersR16ReferencesignalR16ChoiceCsiRsIndexR16
)

type PuschPathlossreferencersR16ReferencesignalR16 struct {
	Choice        uint64
	SsbIndexR16   *SsbIndex
	CsiRsIndexR16 *NzpCsiRsResourceid
}
