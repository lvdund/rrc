package ies

import "rrc/utils"

// PRACH-ConfigInfo ::= SEQUENCE
type PrachConfiginfo struct {
	PrachConfigindex          utils.INTEGER `lb:0,ub:63`
	Highspeedflag             utils.BOOLEAN
	Zerocorrelationzoneconfig utils.INTEGER `lb:0,ub:15`
	PrachFreqoffset           utils.INTEGER `lb:0,ub:94`
}
