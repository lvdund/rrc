package ies

import "rrc/utils"

// PRACH-Config ::= SEQUENCE
type PrachConfig struct {
	Rootsequenceindex utils.INTEGER
	PrachConfiginfo   *PrachConfiginfo
}
