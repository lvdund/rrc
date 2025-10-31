package ies

// RRCConnectionSetupComplete-NB-criticalExtensions ::= CHOICE
const (
	RrcconnectionsetupcompleteNbCriticalextensionsChoiceNothing = iota
	RrcconnectionsetupcompleteNbCriticalextensionsChoiceRrcconnectionsetupcompleteR13
	RrcconnectionsetupcompleteNbCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcconnectionsetupcompleteNbCriticalextensions struct {
	Choice                        uint64
	RrcconnectionsetupcompleteR13 *RrcconnectionsetupcompleteNbR13
	Criticalextensionsfuture      *RrcconnectionsetupcompleteNbCriticalextensionsCriticalextensionsfuture
}
