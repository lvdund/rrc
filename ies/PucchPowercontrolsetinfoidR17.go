package ies

import "rrc/utils"

// PUCCH-PowerControlSetInfoId-r17 ::= utils.INTEGER (1.. maxNrofPowerControlSetInfos-r17)
type PucchPowercontrolsetinfoidR17 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofPowerControlSetInfosR17`
}
