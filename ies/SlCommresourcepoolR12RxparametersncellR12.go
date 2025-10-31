package ies

import "rrc/utils"

// SL-CommResourcePool-r12-rxParametersNCell-r12 ::= SEQUENCE
type SlCommresourcepoolR12RxparametersncellR12 struct {
	TddConfigR12       *TddConfig
	SyncconfigindexR12 utils.INTEGER `lb:0,ub:15`
}
