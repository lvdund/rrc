package ies

import "rrc/utils"

// RACH-ConfigCommon-groupBconfigured ::= SEQUENCE
type RachConfigcommonGroupbconfigured struct {
	RaMsg3sizegroupa          RachConfigcommonGroupbconfiguredRaMsg3sizegroupa
	Messagepoweroffsetgroupb  RachConfigcommonGroupbconfiguredMessagepoweroffsetgroupb
	NumberofraPreamblesgroupa utils.INTEGER `lb:0,ub:64`
}
