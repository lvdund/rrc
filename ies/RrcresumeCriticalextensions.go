package ies

// RRCResume-criticalExtensions ::= CHOICE
const (
	RrcresumeCriticalextensionsChoiceNothing = iota
	RrcresumeCriticalextensionsChoiceRrcresume
	RrcresumeCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcresumeCriticalextensions struct {
	Choice                   uint64
	Rrcresume                *Rrcresume
	Criticalextensionsfuture *RrcresumeCriticalextensionsCriticalextensionsfuture
}
