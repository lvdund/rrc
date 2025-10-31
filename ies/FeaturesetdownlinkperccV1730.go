package ies

import "rrc/utils"

// FeatureSetDownlinkPerCC-v1730 ::= SEQUENCE
type FeaturesetdownlinkperccV1730 struct {
	IntraslottdmUnicastgroupcommonpdschR17 *FeaturesetdownlinkperccV1730IntraslottdmUnicastgroupcommonpdschR17
	SpsMulticastscellR17                   *FeaturesetdownlinkperccV1730SpsMulticastscellR17
	SpsMulticastscellmulticonfigR17        *utils.INTEGER `lb:0,ub:8`
	DciBroadcastwith16repetitionsR17       *FeaturesetdownlinkperccV1730DciBroadcastwith16repetitionsR17
}
