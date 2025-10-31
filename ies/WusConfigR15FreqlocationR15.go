package ies

import "rrc/utils"

// WUS-Config-r15-freqLocation-r15 ::= ENUMERATED
type WusConfigR15FreqlocationR15 struct {
	Value utils.ENUMERATED
}

const (
	WusConfigR15FreqlocationR15EnumeratedNothing = iota
	WusConfigR15FreqlocationR15EnumeratedN0
	WusConfigR15FreqlocationR15EnumeratedN2
	WusConfigR15FreqlocationR15EnumeratedN4
	WusConfigR15FreqlocationR15EnumeratedSpare1
)
