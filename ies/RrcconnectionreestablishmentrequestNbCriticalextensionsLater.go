package ies

// RRCConnectionReestablishmentRequest-NB-criticalExtensions-later ::= CHOICE
const (
	RrcconnectionreestablishmentrequestNbCriticalextensionsLaterChoiceNothing = iota
	RrcconnectionreestablishmentrequestNbCriticalextensionsLaterChoiceRrcconnectionreestablishmentrequestR14
	RrcconnectionreestablishmentrequestNbCriticalextensionsLaterChoiceLater
)

type RrcconnectionreestablishmentrequestNbCriticalextensionsLater struct {
	Choice                                 uint64
	RrcconnectionreestablishmentrequestR14 *RrcconnectionreestablishmentrequestNbR14
	Later                                  *RrcconnectionreestablishmentrequestNbCriticalextensionsLaterLater
}
