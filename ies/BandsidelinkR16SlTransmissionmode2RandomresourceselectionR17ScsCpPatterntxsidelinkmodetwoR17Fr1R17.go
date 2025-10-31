package ies

import "rrc/utils"

// BandSidelink-r16-sl-TransmissionMode2-RandomResourceSelection-r17-scs-CP-PatternTxSidelinkModeTwo-r17-fr1-r17 ::= SEQUENCE
type BandsidelinkR16SlTransmissionmode2RandomresourceselectionR17ScsCpPatterntxsidelinkmodetwoR17Fr1R17 struct {
	Scs15khzR17 *utils.BITSTRING `lb:16,ub:16`
	Scs30khzR17 *utils.BITSTRING `lb:16,ub:16`
	Scs60khzR17 *utils.BITSTRING `lb:16,ub:16`
}
