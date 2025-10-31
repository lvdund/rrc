package ies

import "rrc/utils"

// FeatureSetUplink-v1720-phy-PrioritizationHighPriorityDG-LowPriorityCG-r17 ::= SEQUENCE
type FeaturesetuplinkV1720PhyPrioritizationhighprioritydgLowprioritycgR17 struct {
	PuschPreparationlowpriorityR17 FeaturesetuplinkV1720PhyPrioritizationhighprioritydgLowprioritycgR17PuschPreparationlowpriorityR17
	AdditionalcancellationtimeR17  *FeaturesetuplinkV1720PhyPrioritizationhighprioritydgLowprioritycgR17AdditionalcancellationtimeR17
	MaxnumbercarriersR17           utils.INTEGER `lb:0,ub:16`
}
