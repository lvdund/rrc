package ies

// PURConfigurationRequest-r16-IEs-pur-ConfigRequest-r16 ::= CHOICE
const (
	PurconfigurationrequestR16IesPurConfigrequestR16ChoiceNothing = iota
	PurconfigurationrequestR16IesPurConfigrequestR16ChoicePurReleaserequest
	PurconfigurationrequestR16IesPurConfigrequestR16ChoicePurSetuprequest
)

type PurconfigurationrequestR16IesPurConfigrequestR16 struct {
	Choice            uint64
	PurReleaserequest *struct{}
	PurSetuprequest   *PurconfigurationrequestR16IesPurConfigrequestR16PurSetuprequest
}
