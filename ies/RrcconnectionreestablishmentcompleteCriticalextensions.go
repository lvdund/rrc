package ies

// RRCConnectionReestablishmentComplete-criticalExtensions ::= CHOICE
const (
	RrcconnectionreestablishmentcompleteCriticalextensionsChoiceNothing = iota
	RrcconnectionreestablishmentcompleteCriticalextensionsChoiceRrcconnectionreestablishmentcompleteR8
	RrcconnectionreestablishmentcompleteCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcconnectionreestablishmentcompleteCriticalextensions struct {
	Choice                                 uint64
	RrcconnectionreestablishmentcompleteR8 *RrcconnectionreestablishmentcompleteR8
	Criticalextensionsfuture               *RrcconnectionreestablishmentcompleteCriticalextensionsCriticalextensionsfuture
}
