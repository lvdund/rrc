package ies

import "rrc/utils"

// PUSCH-TPC-CommandConfig ::= SEQUENCE
// Extensible
type PuschTpcCommandconfig struct {
	TpcIndex    *utils.INTEGER `lb:0,ub:15`
	TpcIndexsul *utils.INTEGER `lb:0,ub:15`
	Targetcell  *Servcellindex
}
