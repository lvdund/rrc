package ies

import "rrc/utils"

// SL-CommResourcePoolV2X-r14-rxParametersNCell-r14 ::= SEQUENCE
type SlCommresourcepoolv2xR14RxparametersncellR14 struct {
	TddConfigR14       *TddConfig
	SyncconfigindexR14 utils.INTEGER `lb:0,ub:15`
}
