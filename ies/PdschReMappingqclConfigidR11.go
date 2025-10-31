package ies

import "rrc/utils"

// PDSCH-RE-MappingQCL-ConfigId-r11 ::= utils.INTEGER (1..maxRE-MapQCL-r11)
type PdschReMappingqclConfigidR11 struct {
	Value utils.INTEGER `lb:0,ub:maxREMapqclR11`
}
