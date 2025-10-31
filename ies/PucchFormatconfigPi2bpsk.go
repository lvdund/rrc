package ies

import "rrc/utils"

// PUCCH-FormatConfig-pi2BPSK ::= ENUMERATED
type PucchFormatconfigPi2bpsk struct {
	Value utils.ENUMERATED
}

const (
	PucchFormatconfigPi2bpskEnumeratedNothing = iota
	PucchFormatconfigPi2bpskEnumeratedEnabled
)
