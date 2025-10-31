package ies

// RRCConnectionResumeRequest-NB-criticalExtensions ::= CHOICE
const (
	RrcconnectionresumerequestNbCriticalextensionsChoiceNothing = iota
	RrcconnectionresumerequestNbCriticalextensionsChoiceRrcconnectionresumerequestR13
	RrcconnectionresumerequestNbCriticalextensionsChoiceLater
)

type RrcconnectionresumerequestNbCriticalextensions struct {
	Choice                        uint64
	RrcconnectionresumerequestR13 *RrcconnectionresumerequestNbR13
	Later                         *RrcconnectionresumerequestNbCriticalextensionsLater
}
