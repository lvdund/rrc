package ies

import "rrc/utils"

// PUCCH-TPC-CommandConfig ::= SEQUENCE
// Extensible
type PucchTpcCommandconfig struct {
	TpcIndexpcell                             *utils.INTEGER `lb:0,ub:15`
	TpcIndexpucchScell                        *utils.INTEGER `lb:0,ub:15`
	TpcIndexpucchSscellR17                    *utils.INTEGER `lb:0,ub:15,ext`
	TpcIndexpucchSscellsecondarypucchgroupR17 *utils.INTEGER `lb:0,ub:15,ext`
}
