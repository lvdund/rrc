package ies

import "rrc/utils"

// PRACH-ConfigInfo ::= SEQUENCE
type PrachConfiginfo struct {
	PrachConfigindex          utils.INTEGER
	Highspeedflag             bool
	Zerocorrelationzoneconfig utils.INTEGER
	PrachFreqoffset           utils.INTEGER
}
