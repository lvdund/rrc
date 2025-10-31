package ies

// RRCConnectionReconfigurationComplete-NB-criticalExtensions ::= CHOICE
const (
	RrcconnectionreconfigurationcompleteNbCriticalextensionsChoiceNothing = iota
	RrcconnectionreconfigurationcompleteNbCriticalextensionsChoiceRrcconnectionreconfigurationcompleteR13
	RrcconnectionreconfigurationcompleteNbCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcconnectionreconfigurationcompleteNbCriticalextensions struct {
	Choice                                  uint64
	RrcconnectionreconfigurationcompleteR13 *RrcconnectionreconfigurationcompleteNbR13
	Criticalextensionsfuture                *RrcconnectionreconfigurationcompleteNbCriticalextensionsCriticalextensionsfuture
}
