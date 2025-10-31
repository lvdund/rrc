package ies

// RRCConnectionReconfigurationComplete-criticalExtensions ::= CHOICE
const (
	RrcconnectionreconfigurationcompleteCriticalextensionsChoiceNothing = iota
	RrcconnectionreconfigurationcompleteCriticalextensionsChoiceRrcconnectionreconfigurationcompleteR8
	RrcconnectionreconfigurationcompleteCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcconnectionreconfigurationcompleteCriticalextensions struct {
	Choice                                 uint64
	RrcconnectionreconfigurationcompleteR8 *RrcconnectionreconfigurationcompleteR8
	Criticalextensionsfuture               *RrcconnectionreconfigurationcompleteCriticalextensionsCriticalextensionsfuture
}
