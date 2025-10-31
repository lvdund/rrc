package ies

import "rrc/utils"

// CFR-ConfigMulticast-r17 ::= SEQUENCE
type CfrConfigmulticastR17 struct {
	LocationandbandwidthmulticastR17   *utils.INTEGER `lb:0,ub:37949`
	PdcchConfigmulticastR17            *PdcchConfig
	PdschConfigmulticastR17            *PdschConfig
	SpsConfigmulticasttoaddmodlistR17  *SpsConfigmulticasttoaddmodlistR17
	SpsConfigmulticasttoreleaselistR17 *SpsConfigmulticasttoreleaselistR17
}
