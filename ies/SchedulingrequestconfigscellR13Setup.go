package ies

import "rrc/utils"

// SchedulingRequestConfigSCell-r13-setup ::= SEQUENCE
type SchedulingrequestconfigscellR13Setup struct {
	SrPucchResourceindexR13   utils.INTEGER  `lb:0,ub:2047`
	SrPucchResourceindexp1R13 *utils.INTEGER `lb:0,ub:2047`
	SrConfigindexR13          utils.INTEGER  `lb:0,ub:157`
	DsrTransmaxR13            SchedulingrequestconfigscellR13SetupDsrTransmaxR13
}
