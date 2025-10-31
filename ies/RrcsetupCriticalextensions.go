package ies

// RRCSetup-criticalExtensions ::= CHOICE
const (
	RrcsetupCriticalextensionsChoiceNothing = iota
	RrcsetupCriticalextensionsChoiceRrcsetup
	RrcsetupCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcsetupCriticalextensions struct {
	Choice                   uint64
	Rrcsetup                 *Rrcsetup
	Criticalextensionsfuture *RrcsetupCriticalextensionsCriticalextensionsfuture
}
