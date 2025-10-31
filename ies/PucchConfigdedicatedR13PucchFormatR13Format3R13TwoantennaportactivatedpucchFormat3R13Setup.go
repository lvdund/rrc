package ies

import "rrc/utils"

// PUCCH-ConfigDedicated-r13-pucch-Format-r13-format3-r13-twoAntennaPortActivatedPUCCH-Format3-r13-setup ::= SEQUENCE
type PucchConfigdedicatedR13PucchFormatR13Format3R13TwoantennaportactivatedpucchFormat3R13Setup struct {
	N3pucchAnListp1R13 []utils.INTEGER `lb:1,ub:4`
}
