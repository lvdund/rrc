package ies

// RRCConnectionReestablishment-criticalExtensions ::= CHOICE
const (
	RrcconnectionreestablishmentCriticalextensionsChoiceNothing = iota
	RrcconnectionreestablishmentCriticalextensionsChoiceC1
	RrcconnectionreestablishmentCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcconnectionreestablishmentCriticalextensions struct {
	Choice                   uint64
	C1                       *RrcconnectionreestablishmentCriticalextensionsC1
	Criticalextensionsfuture *RrcconnectionreestablishmentCriticalextensionsCriticalextensionsfuture
}
