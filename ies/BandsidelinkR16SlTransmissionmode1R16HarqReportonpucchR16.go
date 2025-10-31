package ies

import "rrc/utils"

// BandSidelink-r16-sl-TransmissionMode1-r16-harq-ReportOnPUCCH-r16 ::= ENUMERATED
type BandsidelinkR16SlTransmissionmode1R16HarqReportonpucchR16 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16SlTransmissionmode1R16HarqReportonpucchR16EnumeratedNothing = iota
	BandsidelinkR16SlTransmissionmode1R16HarqReportonpucchR16EnumeratedSupported
)
