package ies

import "rrc/utils"

// SRS-TPC-CommandConfig ::= SEQUENCE
// Extensible
type SrsTpcCommandconfig struct {
	Startingbitofformat23    *utils.INTEGER `lb:0,ub:31`
	Fieldtypeformat23        *utils.INTEGER `lb:0,ub:1`
	Startingbitofformat23sul *utils.INTEGER `lb:0,ub:31,ext`
}
