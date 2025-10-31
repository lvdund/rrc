package ies

import "rrc/utils"

// ScramblingId ::= utils.INTEGER (0..1023)
type Scramblingid struct {
	Value utils.INTEGER `lb:0,ub:1023`
}
