package ies

import "rrc/utils"

// SL-DiscResourcePool-r12-rxParameters-r12 ::= SEQUENCE
type SlDiscresourcepoolR12RxparametersR12 struct {
	TddConfigR12       *TddConfig
	SyncconfigindexR12 utils.INTEGER `lb:0,ub:15`
}
