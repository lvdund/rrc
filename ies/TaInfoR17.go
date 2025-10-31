package ies

import "rrc/utils"

// TA-Info-r17 ::= SEQUENCE
type TaInfoR17 struct {
	TaCommonR17             utils.INTEGER  `lb:0,ub:66485757`
	TaCommondriftR17        *utils.INTEGER `lb:0,ub:257303`
	TaCommondriftvariantR17 *utils.INTEGER `lb:0,ub:28949`
}
