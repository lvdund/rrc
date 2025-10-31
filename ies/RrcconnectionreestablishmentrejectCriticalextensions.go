package ies

// RRCConnectionReestablishmentReject-criticalExtensions ::= CHOICE
const (
	RrcconnectionreestablishmentrejectCriticalextensionsChoiceNothing = iota
	RrcconnectionreestablishmentrejectCriticalextensionsChoiceRrcconnectionreestablishmentrejectR8
	RrcconnectionreestablishmentrejectCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcconnectionreestablishmentrejectCriticalextensions struct {
	Choice                               uint64
	RrcconnectionreestablishmentrejectR8 *RrcconnectionreestablishmentrejectR8
	Criticalextensionsfuture             *RrcconnectionreestablishmentrejectCriticalextensionsCriticalextensionsfuture
}
