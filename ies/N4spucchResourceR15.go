package ies

import "rrc/utils"

// N4SPUCCH-Resource-r15 ::= SEQUENCE
type N4spucchResourceR15 struct {
	N4startingprbR15 utils.INTEGER `lb:0,ub:109`
	N4numberofprbR15 utils.INTEGER `lb:0,ub:7`
}
