package ies

import "rrc/utils"

// AppLayerBufferLevel-r17 ::= utils.INTEGER (0..30000)
type ApplayerbufferlevelR17 struct {
	Value utils.INTEGER `lb:0,ub:30000`
}
