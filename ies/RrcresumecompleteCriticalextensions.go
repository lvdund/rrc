package ies

// RRCResumeComplete-criticalExtensions ::= CHOICE
const (
	RrcresumecompleteCriticalextensionsChoiceNothing = iota
	RrcresumecompleteCriticalextensionsChoiceRrcresumecomplete
	RrcresumecompleteCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcresumecompleteCriticalextensions struct {
	Choice                   uint64
	Rrcresumecomplete        *Rrcresumecomplete
	Criticalextensionsfuture *RrcresumecompleteCriticalextensionsCriticalextensionsfuture
}
