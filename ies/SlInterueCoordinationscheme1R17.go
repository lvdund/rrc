package ies

import "rrc/utils"

// SL-InterUE-CoordinationScheme1-r17 ::= SEQUENCE
// Extensible
type SlInterueCoordinationscheme1R17 struct {
	SlIucExplicitR17                          *SlInterueCoordinationscheme1R17SlIucExplicitR17
	SlIucConditionR17                         *SlInterueCoordinationscheme1R17SlIucConditionR17
	SlCondition1A2R17                         *SlInterueCoordinationscheme1R17SlCondition1A2R17
	SlThresholdrsrpCondition1B1Option1listR17 *[]SlThresholdrsrpCondition1B1R17 `lb:1,ub:8`
	SlThresholdrsrpCondition1B1Option2listR17 *[]SlThresholdrsrpCondition1B1R17 `lb:1,ub:8`
	SlContainercoordinfoR17                   *SlInterueCoordinationscheme1R17SlContainercoordinfoR17
	SlContainerrequestR17                     *SlInterueCoordinationscheme1R17SlContainerrequestR17
	SlTriggerconditioncoordinfoR17            *utils.INTEGER `lb:0,ub:1`
	SlTriggerconditionrequestR17              *utils.INTEGER `lb:0,ub:1`
	SlPrioritycoordinfoexplicitR17            *utils.INTEGER `lb:0,ub:8`
	SlPrioritycoordinfoconditionR17           *utils.INTEGER `lb:0,ub:8`
	SlPriorityrequestR17                      *utils.INTEGER `lb:0,ub:8`
	SlPrioritypreferredresourcesetR17         *utils.INTEGER `lb:0,ub:8`
	SlMaxslotoffsettrivR17                    *utils.INTEGER `lb:0,ub:8000`
	SlNumsubchPreferredresourcesetR17         *utils.INTEGER `lb:0,ub:27`
	SlReservedperiodpreferredresourcesetR17   *utils.INTEGER `lb:0,ub:16`
	SlDetermineresourcetypeR17                *SlInterueCoordinationscheme1R17SlDetermineresourcetypeR17
}
