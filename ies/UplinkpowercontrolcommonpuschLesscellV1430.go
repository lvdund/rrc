package ies

import "rrc/utils"

// UplinkPowerControlCommonPUSCH-LessCell-v1430 ::= SEQUENCE
type UplinkpowercontrolcommonpuschLesscellV1430 struct {
	P0NominalPeriodicsrsR14  *utils.INTEGER `lb:0,ub:24`
	P0NominalAperiodicsrsR14 *utils.INTEGER `lb:0,ub:24`
	AlphaSrsR14              *AlphaR12
}
