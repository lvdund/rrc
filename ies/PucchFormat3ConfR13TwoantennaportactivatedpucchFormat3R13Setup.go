package ies

import "rrc/utils"

// PUCCH-Format3-Conf-r13-twoAntennaPortActivatedPUCCH-Format3-r13-setup ::= SEQUENCE
type PucchFormat3ConfR13TwoantennaportactivatedpucchFormat3R13Setup struct {
	N3pucchAnListp1R13 []utils.INTEGER `lb:1,ub:4`
}
