package ies

import "rrc/utils"

// CSI-RS-ConfigNZP-Activation-r14 ::= SEQUENCE
type CsiRsConfignzpActivationR14 struct {
	CsiRsNzpModeR14       utils.ENUMERATED
	ActivatedresourcesR14 utils.INTEGER
}
