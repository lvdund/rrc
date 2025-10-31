package ies

// RRCConnectionResumeRequest-r13-criticalExtensions ::= CHOICE
const (
	RrcconnectionresumerequestR13CriticalextensionsChoiceNothing = iota
	RrcconnectionresumerequestR13CriticalextensionsChoiceRrcconnectionresumerequestR13
	RrcconnectionresumerequestR13CriticalextensionsChoiceRrcconnectionresumerequestR15
)

type RrcconnectionresumerequestR13Criticalextensions struct {
	Choice                        uint64
	RrcconnectionresumerequestR13 *RrcconnectionresumerequestR13
	RrcconnectionresumerequestR15 *Rrcconnectionresumerequest5gcR15
}
