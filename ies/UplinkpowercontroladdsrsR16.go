package ies

import "rrc/utils"

// UplinkPowerControlAddSRS-r16 ::= SEQUENCE
type UplinkpowercontroladdsrsR16 struct {
	TpcIndexsrsAddR16              *TpcIndex
	Startingbitofformat3bSrsAddR16 *utils.INTEGER `lb:0,ub:31`
	Fieldtypeformat3bSrsAddR16     *utils.INTEGER `lb:0,ub:2`
	P0UeSrsAddR16                  *utils.INTEGER `lb:0,ub:15`
	AccumulationenabledsrsAddR16   utils.BOOLEAN
}
