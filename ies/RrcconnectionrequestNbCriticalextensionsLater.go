package ies

// RRCConnectionRequest-NB-criticalExtensions-later ::= CHOICE
const (
	RrcconnectionrequestNbCriticalextensionsLaterChoiceNothing = iota
	RrcconnectionrequestNbCriticalextensionsLaterChoiceRrcconnectionrequestR16
	RrcconnectionrequestNbCriticalextensionsLaterChoiceCriticalextensionsfuture
)

type RrcconnectionrequestNbCriticalextensionsLater struct {
	Choice                   uint64
	RrcconnectionrequestR16  *Rrcconnectionrequest5gcNbR16
	Criticalextensionsfuture *RrcconnectionrequestNbCriticalextensionsLaterCriticalextensionsfuture
}
