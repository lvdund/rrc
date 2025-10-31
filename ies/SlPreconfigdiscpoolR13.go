package ies

import "rrc/utils"

// SL-PreconfigDiscPool-r13 ::= SEQUENCE
// Extensible
type SlPreconfigdiscpoolR13 struct {
	CpLenR13            SlCpLenR12
	DiscperiodR13       SlPreconfigdiscpoolR13DiscperiodR13
	NumretxR13          utils.INTEGER `lb:0,ub:3`
	NumrepetitionR13    utils.INTEGER `lb:0,ub:50`
	TfResourceconfigR13 SlTfResourceconfigR12
	TxparametersR13     *SlPreconfigdiscpoolR13TxparametersR13
}
