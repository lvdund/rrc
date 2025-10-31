package ies

import "rrc/utils"

// PUSCH-Config-ul-FullPowerTransmission-r16 ::= ENUMERATED
type PuschConfigUlFullpowertransmissionR16 struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigUlFullpowertransmissionR16EnumeratedNothing = iota
	PuschConfigUlFullpowertransmissionR16EnumeratedFullpower
	PuschConfigUlFullpowertransmissionR16EnumeratedFullpowermode1
	PuschConfigUlFullpowertransmissionR16EnumeratedFullpowermode2
)
