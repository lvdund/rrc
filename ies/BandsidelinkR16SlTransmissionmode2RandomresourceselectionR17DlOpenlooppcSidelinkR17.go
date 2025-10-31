package ies

import "rrc/utils"

// BandSidelink-r16-sl-TransmissionMode2-RandomResourceSelection-r17-dl-openLoopPC-Sidelink-r17 ::= ENUMERATED
type BandsidelinkR16SlTransmissionmode2RandomresourceselectionR17DlOpenlooppcSidelinkR17 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16SlTransmissionmode2RandomresourceselectionR17DlOpenlooppcSidelinkR17EnumeratedNothing = iota
	BandsidelinkR16SlTransmissionmode2RandomresourceselectionR17DlOpenlooppcSidelinkR17EnumeratedSupported
)
