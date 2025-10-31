package ies

import "rrc/utils"

// PUSCH-PowerControl ::= SEQUENCE
type PuschPowercontrol struct {
	TpcAccumulation                  *PuschPowercontrolTpcAccumulation
	Msg3Alpha                        *Alpha
	P0Nominalwithoutgrant            *utils.INTEGER                `lb:0,ub:24`
	P0Alphasets                      *[]P0PuschAlphaset            `lb:1,ub:maxNrofP0PuschAlphasets`
	Pathlossreferencerstoaddmodlist  *[]PuschPathlossreferencers   `lb:1,ub:maxNrofPUSCHPathlossreferencerss`
	Pathlossreferencerstoreleaselist *[]PuschPathlossreferencersId `lb:1,ub:maxNrofPUSCHPathlossreferencerss`
	TwopuschPcAdjustmentstates       *PuschPowercontrolTwopuschPcAdjustmentstates
	Deltamcs                         *PuschPowercontrolDeltamcs
	SriPuschMappingtoaddmodlist      *[]SriPuschPowercontrol   `lb:1,ub:maxNrofSRIPuschMappings`
	SriPuschMappingtoreleaselist     *[]SriPuschPowercontrolid `lb:1,ub:maxNrofSRIPuschMappings`
}
