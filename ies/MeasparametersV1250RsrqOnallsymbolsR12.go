package ies

import "rrc/utils"

// MeasParameters-v1250-rsrq-OnAllSymbols-r12 ::= ENUMERATED
type MeasparametersV1250RsrqOnallsymbolsR12 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1250RsrqOnallsymbolsR12EnumeratedNothing = iota
	MeasparametersV1250RsrqOnallsymbolsR12EnumeratedSupported
)
