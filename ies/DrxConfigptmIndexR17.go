package ies

import "rrc/utils"

// DRX-ConfigPTM-Index-r17 ::= utils.INTEGER (0..maxNrofDRX-ConfigPTM-1-r17)
type DrxConfigptmIndexR17 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofDRXConfigptm1R17`
}
