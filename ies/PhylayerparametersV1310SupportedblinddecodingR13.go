package ies

import "rrc/utils"

// PhyLayerParameters-v1310-supportedBlindDecoding-r13 ::= SEQUENCE
type PhylayerparametersV1310SupportedblinddecodingR13 struct {
	MaxnumberdecodingR13          *utils.INTEGER `lb:0,ub:32`
	PdcchCandidatereductionsR13   *PhylayerparametersV1310SupportedblinddecodingR13PdcchCandidatereductionsR13
	SkipmonitoringdciFormat01aR13 *PhylayerparametersV1310SupportedblinddecodingR13SkipmonitoringdciFormat01aR13
}
