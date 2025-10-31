package ies

import "rrc/utils"

// IDC-Config-r11-idc-Indication-r11 ::= ENUMERATED
type IdcConfigR11IdcIndicationR11 struct {
	Value utils.ENUMERATED
}

const (
	IdcConfigR11IdcIndicationR11EnumeratedNothing = iota
	IdcConfigR11IdcIndicationR11EnumeratedSetup
)
