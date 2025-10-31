package ies

import "rrc/utils"

// BandSidelink-r16-sl-TransmissionMode2-r16-dl-openLoopPC-Sidelink-r16 ::= ENUMERATED
type BandsidelinkR16SlTransmissionmode2R16DlOpenlooppcSidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16SlTransmissionmode2R16DlOpenlooppcSidelinkR16EnumeratedNothing = iota
	BandsidelinkR16SlTransmissionmode2R16DlOpenlooppcSidelinkR16EnumeratedSupported
)
