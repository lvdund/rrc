package ies

import "rrc/utils"

// PUCCH-PowerControl ::= SEQUENCE
// Extensible
type PucchPowercontrol struct {
	DeltafPucchF0              *utils.INTEGER              `lb:0,ub:15`
	DeltafPucchF1              *utils.INTEGER              `lb:0,ub:15`
	DeltafPucchF2              *utils.INTEGER              `lb:0,ub:15`
	DeltafPucchF3              *utils.INTEGER              `lb:0,ub:15`
	DeltafPucchF4              *utils.INTEGER              `lb:0,ub:15`
	P0Set                      *[]P0Pucch                  `lb:1,ub:maxNrofPUCCHP0Perset`
	Pathlossreferencerss       *[]PucchPathlossreferencers `lb:1,ub:maxNrofPUCCHPathlossreferencerss`
	TwopucchPcAdjustmentstates *PucchPowercontrolTwopucchPcAdjustmentstates
	PathlossreferencerssV1610  *utils.Setuprelease[PathlossreferencerssV1610] `ext`
}
