package ies

import "rrc/utils"

// BandSidelink-r16-sl-TransmissionMode2-RandomResourceSelection-r17-extendedCP-Mode2Random-r17 ::= ENUMERATED
type BandsidelinkR16SlTransmissionmode2RandomresourceselectionR17ExtendedcpMode2randomR17 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16SlTransmissionmode2RandomresourceselectionR17ExtendedcpMode2randomR17EnumeratedNothing = iota
	BandsidelinkR16SlTransmissionmode2RandomresourceselectionR17ExtendedcpMode2randomR17EnumeratedSupported
)
