package ies

import "rrc/utils"

// SL-DiscTxResourceReq-r13 ::= SEQUENCE
type SlDisctxresourcereqR13 struct {
	CarrierfreqdisctxR13 *utils.INTEGER `lb:0,ub:maxFreq`
	DisctxresourcereqR13 utils.INTEGER  `lb:0,ub:63`
}
