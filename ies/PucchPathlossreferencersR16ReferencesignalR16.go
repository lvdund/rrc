package ies

// PUCCH-PathlossReferenceRS-r16-referenceSignal-r16 ::= CHOICE
const (
	PucchPathlossreferencersR16ReferencesignalR16ChoiceNothing = iota
	PucchPathlossreferencersR16ReferencesignalR16ChoiceSsbIndexR16
	PucchPathlossreferencersR16ReferencesignalR16ChoiceCsiRsIndexR16
)

type PucchPathlossreferencersR16ReferencesignalR16 struct {
	Choice        uint64
	SsbIndexR16   *SsbIndex
	CsiRsIndexR16 *NzpCsiRsResourceid
}
