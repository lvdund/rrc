package ies

import "rrc/utils"

// MIMO-ParametersPerBand-unifiedJointTCI-mTRP-InterCell-BM-r17 ::= SEQUENCE
type MimoParametersperbandUnifiedjointtciMtrpIntercellBmR17 struct {
	MaxnumadditionalpciL1RsrpR17       utils.INTEGER `lb:0,ub:7`
	MaxnumssbResourcel1RsrpAcrossccR17 MimoParametersperbandUnifiedjointtciMtrpIntercellBmR17MaxnumssbResourcel1RsrpAcrossccR17
}
