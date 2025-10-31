package ies

import "rrc/utils"

// BandSidelinkPC5-r16-sl-Rx-256QAM-r16 ::= ENUMERATED
type Bandsidelinkpc5R16SlRx256qamR16 struct {
	Value utils.ENUMERATED
}

const (
	Bandsidelinkpc5R16SlRx256qamR16EnumeratedNothing = iota
	Bandsidelinkpc5R16SlRx256qamR16EnumeratedSupported
)
