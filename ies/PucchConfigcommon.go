package ies

import "rrc/utils"

// PUCCH-ConfigCommon ::= SEQUENCE
// Extensible
type PucchConfigcommon struct {
	PucchResourcecommon          *utils.INTEGER `lb:0,ub:15`
	PucchGrouphopping            PucchConfigcommonPucchGrouphopping
	Hoppingid                    *utils.INTEGER                           `lb:0,ub:1023`
	P0Nominal                    *utils.INTEGER                           `lb:0,ub:24`
	Nrofprbs                     *utils.INTEGER                           `lb:0,ub:16,ext`
	IntraSlotfhR17               *PucchConfigcommonIntraSlotfhR17         `ext`
	PucchResourcecommonredcapR17 *utils.INTEGER                           `lb:0,ub:15,ext`
	AdditionalprboffsetR17       *PucchConfigcommonAdditionalprboffsetR17 `ext`
}
