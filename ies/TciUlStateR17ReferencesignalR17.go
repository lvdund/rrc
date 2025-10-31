package ies

// TCI-UL-State-r17-referenceSignal-r17 ::= CHOICE
const (
	TciUlStateR17ReferencesignalR17ChoiceNothing = iota
	TciUlStateR17ReferencesignalR17ChoiceSsbIndexR17
	TciUlStateR17ReferencesignalR17ChoiceCsiRsIndexR17
	TciUlStateR17ReferencesignalR17ChoiceSrsR17
)

type TciUlStateR17ReferencesignalR17 struct {
	Choice        uint64
	SsbIndexR17   *SsbIndex
	CsiRsIndexR17 *NzpCsiRsResourceid
	SrsR17        *SrsResourceid
}
