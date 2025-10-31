package ies

import "rrc/utils"

// UplinkTxDirectCurrentTwoCarrierInfo-r16 ::= SEQUENCE
type UplinktxdirectcurrenttwocarrierinfoR16 struct {
	ReferencecarrierindexR16   Servcellindex
	Shift7dot5khzR16           utils.BOOLEAN
	TxdirectcurrentlocationR16 utils.INTEGER `lb:0,ub:3301`
}
