package ies

import "rrc/utils"

// CGI-InfoNR-r15-noSIB1-r15 ::= SEQUENCE
type CgiInfonrR15Nosib1R15 struct {
	SsbSubcarrieroffsetR15 utils.INTEGER `lb:0,ub:15`
	PdcchConfigsib1R15     utils.INTEGER `lb:0,ub:255`
}
