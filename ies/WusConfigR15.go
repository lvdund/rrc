package ies

import "rrc/utils"

// WUS-Config-r15 ::= SEQUENCE
type WusConfigR15 struct {
	MaxdurationfactorR15   utils.ENUMERATED
	NumposR15              utils.ENUMERATED
	FreqlocationR15        utils.ENUMERATED
	TimeoffsetdrxR15       utils.ENUMERATED
	TimeoffsetEdrxShortR15 utils.ENUMERATED
	TimeoffsetEdrxLongR15  *utils.ENUMERATED
}
