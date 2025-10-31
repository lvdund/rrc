package ies

import "rrc/utils"

// SPS-PUCCH-AN-r16 ::= SEQUENCE
type SpsPucchAnR16 struct {
	SpsPucchAnResourceidR16 PucchResourceid
	MaxpayloadsizeR16       *utils.INTEGER `lb:0,ub:256`
}
