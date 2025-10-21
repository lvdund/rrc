package ies

import "rrc/utils"

// UplinkPowerControlAddSRS-r16 ::= SEQUENCE
type UplinkpowercontroladdsrsR16 struct {
	TpcIndexsrsAddR16              *TpcIndex
	Startingbitofformat3bSrsAddR16 *utils.INTEGER
	Fieldtypeformat3bSrsAddR16     *utils.INTEGER
	P0UeSrsAddR16                  *utils.INTEGER
	AccumulationenabledsrsAddR16   bool
}
