package ies

// RRCConnectionReestablishment-NB-criticalExtensions ::= CHOICE
const (
	RrcconnectionreestablishmentNbCriticalextensionsChoiceNothing = iota
	RrcconnectionreestablishmentNbCriticalextensionsChoiceC1
	RrcconnectionreestablishmentNbCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcconnectionreestablishmentNbCriticalextensions struct {
	Choice                   uint64
	C1                       *RrcconnectionreestablishmentNbCriticalextensionsC1
	Criticalextensionsfuture *RrcconnectionreestablishmentNbCriticalextensionsCriticalextensionsfuture
}
