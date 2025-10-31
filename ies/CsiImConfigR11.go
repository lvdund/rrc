package ies

import "rrc/utils"

// CSI-IM-Config-r11 ::= SEQUENCE
// Extensible
type CsiImConfigR11 struct {
	CsiImConfigidR11  CsiImConfigidR11
	ResourceconfigR11 utils.INTEGER `lb:0,ub:31`
	SubframeconfigR11 utils.INTEGER `lb:0,ub:154`
}
