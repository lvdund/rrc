package ies

import "rrc/utils"

// SchedulingRequestConfig-v1530-setup ::= SEQUENCE
type SchedulingrequestconfigV1530Setup struct {
	SrSlotspucchIndexfhR15         *utils.INTEGER `lb:0,ub:1319`
	SrSlotspucchIndexnofhR15       *utils.INTEGER `lb:0,ub:3959`
	SrSubslotspucchResourcelistR15 *SrSubslotspucchResourcelistR15
	SrConfigindexslotR15           *utils.INTEGER `lb:0,ub:36`
	SrConfigindexsubslotR15        *utils.INTEGER `lb:0,ub:122`
	DssrTransmaxR15                SchedulingrequestconfigV1530SetupDssrTransmaxR15
}
