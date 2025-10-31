package ies

// RRCConnectionResumeRequest-NB-criticalExtensions-later ::= CHOICE
const (
	RrcconnectionresumerequestNbCriticalextensionsLaterChoiceNothing = iota
	RrcconnectionresumerequestNbCriticalextensionsLaterChoiceRrcconnectionresumerequestR16
	RrcconnectionresumerequestNbCriticalextensionsLaterChoiceCriticalextensionsfuture
)

type RrcconnectionresumerequestNbCriticalextensionsLater struct {
	Choice                        uint64
	RrcconnectionresumerequestR16 *Rrcconnectionresumerequest5gcNbR16
	Criticalextensionsfuture      *RrcconnectionresumerequestNbCriticalextensionsLaterCriticalextensionsfuture
}
