package ies

import "rrc/utils"

// BandSidelink-r16-sl-TransmissionMode2-r16-harq-TxProcessModeTwoSidelink-r16 ::= ENUMERATED
type BandsidelinkR16SlTransmissionmode2R16HarqTxprocessmodetwosidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16SlTransmissionmode2R16HarqTxprocessmodetwosidelinkR16EnumeratedNothing = iota
	BandsidelinkR16SlTransmissionmode2R16HarqTxprocessmodetwosidelinkR16EnumeratedN8
	BandsidelinkR16SlTransmissionmode2R16HarqTxprocessmodetwosidelinkR16EnumeratedN16
)
