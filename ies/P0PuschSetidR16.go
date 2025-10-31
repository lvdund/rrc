package ies

import "rrc/utils"

// P0-PUSCH-SetId-r16 ::= utils.INTEGER (0..maxNrofSRI-PUSCH-Mappings-1)
type P0PuschSetidR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofSRIPuschMappings1`
}
