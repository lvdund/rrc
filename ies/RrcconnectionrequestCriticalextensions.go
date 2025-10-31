package ies

// RRCConnectionRequest-criticalExtensions ::= CHOICE
const (
	RrcconnectionrequestCriticalextensionsChoiceNothing = iota
	RrcconnectionrequestCriticalextensionsChoiceRrcconnectionrequestR8
	RrcconnectionrequestCriticalextensionsChoiceRrcconnectionrequestR15
)

type RrcconnectionrequestCriticalextensions struct {
	Choice                  uint64
	RrcconnectionrequestR8  *RrcconnectionrequestR8
	RrcconnectionrequestR15 *Rrcconnectionrequest5gcR15
}
