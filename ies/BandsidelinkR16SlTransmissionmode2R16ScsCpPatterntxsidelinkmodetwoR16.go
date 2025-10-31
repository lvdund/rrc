package ies

import "rrc/utils"

// BandSidelink-r16-sl-TransmissionMode2-r16-scs-CP-PatternTxSidelinkModeTwo-r16 ::= ENUMERATED
type BandsidelinkR16SlTransmissionmode2R16ScsCpPatterntxsidelinkmodetwoR16 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16SlTransmissionmode2R16ScsCpPatterntxsidelinkmodetwoR16EnumeratedNothing = iota
	BandsidelinkR16SlTransmissionmode2R16ScsCpPatterntxsidelinkmodetwoR16EnumeratedSupported
)
