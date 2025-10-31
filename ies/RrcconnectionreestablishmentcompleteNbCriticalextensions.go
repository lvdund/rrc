package ies

// RRCConnectionReestablishmentComplete-NB-criticalExtensions ::= CHOICE
const (
	RrcconnectionreestablishmentcompleteNbCriticalextensionsChoiceNothing = iota
	RrcconnectionreestablishmentcompleteNbCriticalextensionsChoiceRrcconnectionreestablishmentcompleteR13
	RrcconnectionreestablishmentcompleteNbCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcconnectionreestablishmentcompleteNbCriticalextensions struct {
	Choice                                  uint64
	RrcconnectionreestablishmentcompleteR13 *RrcconnectionreestablishmentcompleteNbR13
	Criticalextensionsfuture                *RrcconnectionreestablishmentcompleteNbCriticalextensionsCriticalextensionsfuture
}
