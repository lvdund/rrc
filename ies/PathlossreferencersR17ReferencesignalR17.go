package ies

// PathlossReferenceRS-r17-referenceSignal-r17 ::= CHOICE
const (
	PathlossreferencersR17ReferencesignalR17ChoiceNothing = iota
	PathlossreferencersR17ReferencesignalR17ChoiceSsbIndex
	PathlossreferencersR17ReferencesignalR17ChoiceCsiRsIndex
)

type PathlossreferencersR17ReferencesignalR17 struct {
	Choice     uint64
	SsbIndex   *SsbIndex
	CsiRsIndex *NzpCsiRsResourceid
}
