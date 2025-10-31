package ies

import "rrc/utils"

// EpochTime-r17 ::= SEQUENCE
type EpochtimeR17 struct {
	SfnR17        utils.INTEGER `lb:0,ub:1023`
	SubframenrR17 utils.INTEGER `lb:0,ub:9`
}
