package ies

// RRCConnectionRequest-NB-criticalExtensions ::= CHOICE
const (
	RrcconnectionrequestNbCriticalextensionsChoiceNothing = iota
	RrcconnectionrequestNbCriticalextensionsChoiceRrcconnectionrequestR13
	RrcconnectionrequestNbCriticalextensionsChoiceLater
)

type RrcconnectionrequestNbCriticalextensions struct {
	Choice                  uint64
	RrcconnectionrequestR13 *RrcconnectionrequestNbR13
	Later                   *RrcconnectionrequestNbCriticalextensionsLater
}
