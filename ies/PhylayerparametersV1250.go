package ies

import "rrc/utils"

// PhyLayerParameters-v1250 ::= SEQUENCE
type PhylayerparametersV1250 struct {
	EHarqPatternFddR12                     *PhylayerparametersV1250EHarqPatternFddR12
	Enhanced4txcodebookR12                 *PhylayerparametersV1250Enhanced4txcodebookR12
	TddFddCaPcellduplexR12                 *utils.BITSTRING `lb:2,ub:2`
	PhyTddReconfigTddPcellR12              *PhylayerparametersV1250PhyTddReconfigTddPcellR12
	PhyTddReconfigFddPcellR12              *PhylayerparametersV1250PhyTddReconfigFddPcellR12
	PuschFeedbackmodeR12                   *PhylayerparametersV1250PuschFeedbackmodeR12
	PuschSrsPowercontrolSubframesetR12     *PhylayerparametersV1250PuschSrsPowercontrolSubframesetR12
	CsiSubframesetR12                      *PhylayerparametersV1250CsiSubframesetR12
	NoresourcerestrictionforttibundlingR12 *PhylayerparametersV1250NoresourcerestrictionforttibundlingR12
	DiscoverysignalsindeactscellR12        *PhylayerparametersV1250DiscoverysignalsindeactscellR12
	NaicsCapabilityListR12                 *NaicsCapabilityListR12
}
