package ies

import "rrc/utils"

// SL-HoppingConfigComm-r12 ::= SEQUENCE
type SlHoppingconfigcommR12 struct {
	HoppingparameterR12 utils.INTEGER `lb:0,ub:504`
	NumsubbandsR12      SlHoppingconfigcommR12NumsubbandsR12
	RbOffsetR12         utils.INTEGER `lb:0,ub:110`
}
