package ies

// RRCReject-criticalExtensions ::= CHOICE
const (
	RrcrejectCriticalextensionsChoiceNothing = iota
	RrcrejectCriticalextensionsChoiceRrcreject
	RrcrejectCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcrejectCriticalextensions struct {
	Choice                   uint64
	Rrcreject                *Rrcreject
	Criticalextensionsfuture *RrcrejectCriticalextensionsCriticalextensionsfuture
}
