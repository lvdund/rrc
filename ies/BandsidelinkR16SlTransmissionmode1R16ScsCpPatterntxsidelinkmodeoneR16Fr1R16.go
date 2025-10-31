package ies

import "rrc/utils"

// BandSidelink-r16-sl-TransmissionMode1-r16-scs-CP-PatternTxSidelinkModeOne-r16-fr1-r16 ::= SEQUENCE
type BandsidelinkR16SlTransmissionmode1R16ScsCpPatterntxsidelinkmodeoneR16Fr1R16 struct {
	Scs15khzR16 *utils.BITSTRING `lb:16,ub:16`
	Scs30khzR16 *utils.BITSTRING `lb:16,ub:16`
	Scs60khzR16 *utils.BITSTRING `lb:16,ub:16`
}
