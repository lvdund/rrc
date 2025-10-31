package ies

import "rrc/utils"

// BandSidelink-r16-sl-TransmissionMode1-r16-extendedCP-TxSidelink-r16 ::= ENUMERATED
type BandsidelinkR16SlTransmissionmode1R16ExtendedcpTxsidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16SlTransmissionmode1R16ExtendedcpTxsidelinkR16EnumeratedNothing = iota
	BandsidelinkR16SlTransmissionmode1R16ExtendedcpTxsidelinkR16EnumeratedSupported
)
