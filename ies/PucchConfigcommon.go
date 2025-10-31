package ies

import "rrc/utils"

// PUCCH-ConfigCommon ::= SEQUENCE
type PucchConfigcommon struct {
	DeltapucchShift PucchConfigcommonDeltapucchShift
	NrbCqi          utils.INTEGER `lb:0,ub:98`
	NcsAn           utils.INTEGER `lb:0,ub:7`
	N1pucchAn       utils.INTEGER `lb:0,ub:2047`
}
