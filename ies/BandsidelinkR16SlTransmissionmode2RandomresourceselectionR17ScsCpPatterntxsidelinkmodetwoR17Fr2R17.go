package ies

import "rrc/utils"

// BandSidelink-r16-sl-TransmissionMode2-RandomResourceSelection-r17-scs-CP-PatternTxSidelinkModeTwo-r17-fr2-r17 ::= SEQUENCE
type BandsidelinkR16SlTransmissionmode2RandomresourceselectionR17ScsCpPatterntxsidelinkmodetwoR17Fr2R17 struct {
	Scs60khzR17  *utils.BITSTRING `lb:16,ub:16`
	Scs120khzR17 *utils.BITSTRING `lb:16,ub:16`
}
