package ies

import "rrc/utils"

// CrossCarrierSchedulingSCell-SpCell-r17-supportedSCS-Combinations-r17 ::= SEQUENCE
type CrosscarrierschedulingscellSpcellR17SupportedscsCombinationsR17 struct {
	Scs15khz15khzR17 *CrosscarrierschedulingscellSpcellR17SupportedscsCombinationsR17Scs15khz15khzR17
	Scs15khz30khzR17 *CrosscarrierschedulingscellSpcellR17SupportedscsCombinationsR17Scs15khz30khzR17
	Scs15khz60khzR17 *CrosscarrierschedulingscellSpcellR17SupportedscsCombinationsR17Scs15khz60khzR17
	Scs30khz30khzR17 *utils.BITSTRING `lb:1,ub:496`
	Scs30khz60khzR17 *utils.BITSTRING `lb:1,ub:496`
	Scs60khz60khzR17 *utils.BITSTRING `lb:1,ub:496`
}
