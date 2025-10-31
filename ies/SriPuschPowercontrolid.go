package ies

import "rrc/utils"

// SRI-PUSCH-PowerControlId ::= utils.INTEGER (0..maxNrofSRI-PUSCH-Mappings-1)
type SriPuschPowercontrolid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofSRIPuschMappings1`
}
