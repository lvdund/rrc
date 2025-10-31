package ies

import "rrc/utils"

// BandSidelink-r16-sl-Tx-256QAM-r16 ::= ENUMERATED
type BandsidelinkR16SlTx256qamR16 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16SlTx256qamR16EnumeratedNothing = iota
	BandsidelinkR16SlTx256qamR16EnumeratedSupported
)
