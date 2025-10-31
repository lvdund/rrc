package ies

import "rrc/utils"

// PRACH-Config ::= SEQUENCE
type PrachConfig struct {
	Rootsequenceindex utils.INTEGER `lb:0,ub:837`
	PrachConfiginfo   *PrachConfiginfo
}
