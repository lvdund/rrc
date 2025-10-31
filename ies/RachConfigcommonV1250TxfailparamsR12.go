package ies

import "rrc/utils"

// RACH-ConfigCommon-v1250-txFailParams-r12 ::= SEQUENCE
type RachConfigcommonV1250TxfailparamsR12 struct {
	ConnestfailcountR12          RachConfigcommonV1250TxfailparamsR12ConnestfailcountR12
	ConnestfailoffsetvalidityR12 RachConfigcommonV1250TxfailparamsR12ConnestfailoffsetvalidityR12
	ConnestfailoffsetR12         *utils.INTEGER `lb:0,ub:15`
}
