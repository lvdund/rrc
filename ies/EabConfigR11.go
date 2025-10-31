package ies

import "rrc/utils"

// EAB-Config-r11 ::= SEQUENCE
type EabConfigR11 struct {
	EabCategoryR11      EabConfigR11EabCategoryR11
	EabBarringbitmapR11 utils.BITSTRING `lb:10,ub:10`
}
