package ies

import "rrc/utils"

// BandSidelink-r16-sl-TransmissionMode1-r16-harq-TxProcessModeOneSidelink-r16 ::= ENUMERATED
type BandsidelinkR16SlTransmissionmode1R16HarqTxprocessmodeonesidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16SlTransmissionmode1R16HarqTxprocessmodeonesidelinkR16EnumeratedNothing = iota
	BandsidelinkR16SlTransmissionmode1R16HarqTxprocessmodeonesidelinkR16EnumeratedN8
	BandsidelinkR16SlTransmissionmode1R16HarqTxprocessmodeonesidelinkR16EnumeratedN16
)
