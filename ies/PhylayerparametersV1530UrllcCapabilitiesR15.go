package ies

import "rrc/utils"

// PhyLayerParameters-v1530-urllc-Capabilities-r15 ::= SEQUENCE
type PhylayerparametersV1530UrllcCapabilitiesR15 struct {
	PdschRepsubframeR15            *PhylayerparametersV1530UrllcCapabilitiesR15PdschRepsubframeR15
	PdschRepslotR15                *PhylayerparametersV1530UrllcCapabilitiesR15PdschRepslotR15
	PdschRepsubslotR15             *PhylayerparametersV1530UrllcCapabilitiesR15PdschRepsubslotR15
	PuschSpsMulticonfigsubframeR15 *utils.INTEGER `lb:0,ub:6`
	PuschSpsMaxconfigsubframeR15   *utils.INTEGER `lb:0,ub:31`
	PuschSpsMulticonfigslotR15     *utils.INTEGER `lb:0,ub:6`
	PuschSpsMaxconfigslotR15       *utils.INTEGER `lb:0,ub:31`
	PuschSpsMulticonfigsubslotR15  *utils.INTEGER `lb:0,ub:6`
	PuschSpsMaxconfigsubslotR15    *utils.INTEGER `lb:0,ub:31`
	PuschSpsSlotreppcellR15        *PhylayerparametersV1530UrllcCapabilitiesR15PuschSpsSlotreppcellR15
	PuschSpsSlotreppscellR15       *PhylayerparametersV1530UrllcCapabilitiesR15PuschSpsSlotreppscellR15
	PuschSpsSlotrepscellR15        *PhylayerparametersV1530UrllcCapabilitiesR15PuschSpsSlotrepscellR15
	PuschSpsSubframereppcellR15    *PhylayerparametersV1530UrllcCapabilitiesR15PuschSpsSubframereppcellR15
	PuschSpsSubframereppscellR15   *PhylayerparametersV1530UrllcCapabilitiesR15PuschSpsSubframereppscellR15
	PuschSpsSubframerepscellR15    *PhylayerparametersV1530UrllcCapabilitiesR15PuschSpsSubframerepscellR15
	PuschSpsSubslotreppcellR15     *PhylayerparametersV1530UrllcCapabilitiesR15PuschSpsSubslotreppcellR15
	PuschSpsSubslotreppscellR15    *PhylayerparametersV1530UrllcCapabilitiesR15PuschSpsSubslotreppscellR15
	PuschSpsSubslotrepscellR15     *PhylayerparametersV1530UrllcCapabilitiesR15PuschSpsSubslotrepscellR15
	SemistaticcfiR15               *PhylayerparametersV1530UrllcCapabilitiesR15SemistaticcfiR15
	SemistaticcfiPatternR15        *PhylayerparametersV1530UrllcCapabilitiesR15SemistaticcfiPatternR15
}
