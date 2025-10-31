package ies

// RRCConnectionReestablishmentRequest-NB-criticalExtensions ::= CHOICE
const (
	RrcconnectionreestablishmentrequestNbCriticalextensionsChoiceNothing = iota
	RrcconnectionreestablishmentrequestNbCriticalextensionsChoiceRrcconnectionreestablishmentrequestR13
	RrcconnectionreestablishmentrequestNbCriticalextensionsChoiceLater
)

type RrcconnectionreestablishmentrequestNbCriticalextensions struct {
	Choice                                 uint64
	RrcconnectionreestablishmentrequestR13 *RrcconnectionreestablishmentrequestNbR13
	Later                                  *RrcconnectionreestablishmentrequestNbCriticalextensionsLater
}
