package ies

import "rrc/utils"

// CGI-InfoNR-noSIB1 ::= SEQUENCE
type CgiInfonrNosib1 struct {
	SsbSubcarrieroffset utils.INTEGER `lb:0,ub:15`
	PdcchConfigsib1     PdcchConfigsib1
}
