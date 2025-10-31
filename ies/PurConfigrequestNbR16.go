package ies

// PUR-ConfigRequest-NB-r16 ::= CHOICE
const (
	PurConfigrequestNbR16ChoiceNothing = iota
	PurConfigrequestNbR16ChoicePurReleaserequest
	PurConfigrequestNbR16ChoicePurSetuprequest
)

type PurConfigrequestNbR16 struct {
	Choice            uint64
	PurReleaserequest *struct{}
	PurSetuprequest   *PurConfigrequestNbR16PurSetuprequest
}
