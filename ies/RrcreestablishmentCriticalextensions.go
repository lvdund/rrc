package ies

// RRCReestablishment-criticalExtensions ::= CHOICE
const (
	RrcreestablishmentCriticalextensionsChoiceNothing = iota
	RrcreestablishmentCriticalextensionsChoiceRrcreestablishment
	RrcreestablishmentCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcreestablishmentCriticalextensions struct {
	Choice                   uint64
	Rrcreestablishment       *Rrcreestablishment
	Criticalextensionsfuture *RrcreestablishmentCriticalextensionsCriticalextensionsfuture
}
