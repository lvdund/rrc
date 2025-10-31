package ies

import "rrc/utils"

// TimeSinceCHO-Reconfig-r17 ::= utils.INTEGER (0..1023)
type TimesincechoReconfigR17 struct {
	Value utils.INTEGER `lb:0,ub:1023`
}
