package ies

import "rrc/utils"

// MIMO-ParametersPerBand-supportInter-slotTDM-r16 ::= SEQUENCE
type MimoParametersperbandSupportinterSlottdmR16 struct {
	SupportrepnumpdschTdraR16 MimoParametersperbandSupportinterSlottdmR16SupportrepnumpdschTdraR16
	MaxtbsSizeR16             MimoParametersperbandSupportinterSlottdmR16MaxtbsSizeR16
	MaxnumbertciStatesR16     utils.INTEGER `lb:0,ub:2`
}
