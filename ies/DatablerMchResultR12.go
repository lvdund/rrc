package ies

import "rrc/utils"

// DataBLER-MCH-Result-r12 ::= SEQUENCE
type DatablerMchResultR12 struct {
	MchIndexR12       utils.INTEGER `lb:0,ub:maxPMCHPermbsfn`
	DatablerResultR12 BlerResultR12
}
