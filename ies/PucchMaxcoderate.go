package ies

import "rrc/utils"

// PUCCH-MaxCodeRate ::= ENUMERATED
type PucchMaxcoderate struct {
	Value utils.ENUMERATED
}

const (
	PucchMaxcoderateEnumeratedNothing = iota
	PucchMaxcoderateEnumeratedZerodot08
	PucchMaxcoderateEnumeratedZerodot15
	PucchMaxcoderateEnumeratedZerodot25
	PucchMaxcoderateEnumeratedZerodot35
	PucchMaxcoderateEnumeratedZerodot45
	PucchMaxcoderateEnumeratedZerodot60
	PucchMaxcoderateEnumeratedZerodot80
)
