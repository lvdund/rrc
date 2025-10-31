package ies

// MPE-Resource-r17-mpe-ReferenceSignal-r17 ::= CHOICE
const (
	MpeResourceR17MpeReferencesignalR17ChoiceNothing = iota
	MpeResourceR17MpeReferencesignalR17ChoiceCsiRsResourceR17
	MpeResourceR17MpeReferencesignalR17ChoiceSsbResourceR17
)

type MpeResourceR17MpeReferencesignalR17 struct {
	Choice           uint64
	CsiRsResourceR17 *NzpCsiRsResourceid
	SsbResourceR17   *SsbIndex
}
