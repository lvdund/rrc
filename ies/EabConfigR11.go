package ies

import "rrc/utils"

// EAB-Config-r11 ::= SEQUENCE
type EabConfigR11 struct {
	EabCategoryR11      utils.ENUMERATED
	EabBarringbitmapR11 utils.BITSTRING
}
