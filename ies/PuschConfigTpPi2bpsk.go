package ies

import "rrc/utils"

// PUSCH-Config-tp-pi2BPSK ::= ENUMERATED
type PuschConfigTpPi2bpsk struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigTpPi2bpskEnumeratedNothing = iota
	PuschConfigTpPi2bpskEnumeratedEnabled
)
