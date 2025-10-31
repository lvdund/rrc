package ies

import "rrc/utils"

// MeasResultRLFNR-r16-measResult-r16-rsIndexResults-r16 ::= SEQUENCE
type MeasresultrlfnrR16MeasresultR16RsindexresultsR16 struct {
	ResultsssbIndexesR16    *ResultsperssbIndexlist
	SsbrlmconfigbitmapR16   *utils.BITSTRING `lb:64,ub:64`
	ResultscsiRsIndexesR16  *ResultspercsiRsIndexlist
	CsiRsrlmconfigbitmapR16 *utils.BITSTRING `lb:96,ub:96`
}
