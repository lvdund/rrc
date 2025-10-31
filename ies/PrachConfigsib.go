package ies

import "rrc/utils"

// PRACH-ConfigSIB ::= SEQUENCE
type PrachConfigsib struct {
	Rootsequenceindex utils.INTEGER `lb:0,ub:837`
	PrachConfiginfo   PrachConfiginfo
}
