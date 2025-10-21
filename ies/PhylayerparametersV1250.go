package ies

import "rrc/utils"

// PhyLayerParameters-v1250 ::= SEQUENCE
type PhylayerparametersV1250 struct {
	EHarqPatternFddR12                     *utils.ENUMERATED
	Enhanced4txcodebookR12                 *utils.ENUMERATED
	TddFddCaPcellduplexR12                 *utils.BITSTRING
	PhyTddReconfigTddPcellR12              *utils.ENUMERATED
	PhyTddReconfigFddPcellR12              *utils.ENUMERATED
	PuschFeedbackmodeR12                   *utils.ENUMERATED
	PuschSrsPowercontrolSubframesetR12     *utils.ENUMERATED
	CsiSubframesetR12                      *utils.ENUMERATED
	NoresourcerestrictionforttibundlingR12 *utils.ENUMERATED
	DiscoverysignalsindeactscellR12        *utils.ENUMERATED
	NaicsCapabilityListR12                 *NaicsCapabilityListR12
}
