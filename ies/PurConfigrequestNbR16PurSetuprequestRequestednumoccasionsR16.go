package ies

import "rrc/utils"

// PUR-ConfigRequest-NB-r16-pur-SetupRequest-requestedNumOccasions-r16 ::= ENUMERATED
type PurConfigrequestNbR16PurSetuprequestRequestednumoccasionsR16 struct {
	Value utils.ENUMERATED
}

const (
	PurConfigrequestNbR16PurSetuprequestRequestednumoccasionsR16EnumeratedNothing = iota
	PurConfigrequestNbR16PurSetuprequestRequestednumoccasionsR16EnumeratedOne
	PurConfigrequestNbR16PurSetuprequestRequestednumoccasionsR16EnumeratedInfinite
)
