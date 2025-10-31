package ies

// RRCConnectionReestablishmentRequest-criticalExtensions ::= CHOICE
const (
	RrcconnectionreestablishmentrequestCriticalextensionsChoiceNothing = iota
	RrcconnectionreestablishmentrequestCriticalextensionsChoiceRrcconnectionreestablishmentrequestR8
	RrcconnectionreestablishmentrequestCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcconnectionreestablishmentrequestCriticalextensions struct {
	Choice                                uint64
	RrcconnectionreestablishmentrequestR8 *RrcconnectionreestablishmentrequestR8
	Criticalextensionsfuture              *RrcconnectionreestablishmentrequestCriticalextensionsCriticalextensionsfuture
}
