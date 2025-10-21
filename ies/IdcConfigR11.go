package ies

import "rrc/utils"

// IDC-Config-r11 ::= SEQUENCE
// Extensible
type IdcConfigR11 struct {
	IdcIndicationR11              *utils.ENUMERATED
	AutonomousdenialparametersR11 *interface{}
}
