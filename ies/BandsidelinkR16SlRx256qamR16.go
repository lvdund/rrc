package ies

import "rrc/utils"

// BandSidelink-r16-sl-Rx-256QAM-r16 ::= ENUMERATED
type BandsidelinkR16SlRx256qamR16 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16SlRx256qamR16EnumeratedNothing = iota
	BandsidelinkR16SlRx256qamR16EnumeratedSupported
)
