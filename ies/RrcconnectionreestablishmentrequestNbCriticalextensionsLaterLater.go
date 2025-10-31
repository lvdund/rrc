package ies

// RRCConnectionReestablishmentRequest-NB-criticalExtensions-later-later ::= CHOICE
const (
	RrcconnectionreestablishmentrequestNbCriticalextensionsLaterLaterChoiceNothing = iota
	RrcconnectionreestablishmentrequestNbCriticalextensionsLaterLaterChoiceRrcconnectionreestablishmentrequestR16
	RrcconnectionreestablishmentrequestNbCriticalextensionsLaterLaterChoiceCriticalextensionsfuture
)

type RrcconnectionreestablishmentrequestNbCriticalextensionsLaterLater struct {
	Choice                                 uint64
	RrcconnectionreestablishmentrequestR16 *Rrcconnectionreestablishmentrequest5gcNbR16
	Criticalextensionsfuture               *RrcconnectionreestablishmentrequestNbCriticalextensionsLaterLaterCriticalextensionsfuture
}
