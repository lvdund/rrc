package ies

import "rrc/utils"

// SchedulingRequestConfig-setup ::= SEQUENCE
type SchedulingrequestconfigSetup struct {
	SrPucchResourceindex utils.INTEGER `lb:0,ub:2047`
	SrConfigindex        utils.INTEGER `lb:0,ub:157`
	DsrTransmax          SchedulingrequestconfigSetupDsrTransmax
}
