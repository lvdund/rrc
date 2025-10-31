package ies

import "rrc/utils"

// CSI-RS-ConfigNZP-Activation-r14 ::= SEQUENCE
type CsiRsConfignzpActivationR14 struct {
	CsiRsNzpModeR14       CsiRsConfignzpActivationR14CsiRsNzpModeR14
	ActivatedresourcesR14 utils.INTEGER `lb:0,ub:4`
}
