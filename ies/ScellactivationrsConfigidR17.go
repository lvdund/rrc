package ies

import "rrc/utils"

// SCellActivationRS-ConfigId-r17 ::= utils.INTEGER (1.. maxNrofSCellActRS-r17)
type ScellactivationrsConfigidR17 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofSCellActRSR17`
}
