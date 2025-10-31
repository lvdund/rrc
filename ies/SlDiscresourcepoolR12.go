package ies

import "rrc/utils"

// SL-DiscResourcePool-r12 ::= SEQUENCE
// Extensible
type SlDiscresourcepoolR12 struct {
	CpLenR12            SlCpLenR12
	DiscperiodR12       SlDiscresourcepoolR12DiscperiodR12
	NumretxR12          utils.INTEGER `lb:0,ub:3`
	NumrepetitionR12    utils.INTEGER `lb:0,ub:50`
	TfResourceconfigR12 SlTfResourceconfigR12
	TxparametersR12     *SlDiscresourcepoolR12TxparametersR12
	RxparametersR12     *SlDiscresourcepoolR12RxparametersR12
}
