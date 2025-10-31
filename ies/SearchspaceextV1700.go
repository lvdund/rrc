package ies

import "rrc/utils"

// SearchSpaceExt-v1700 ::= SEQUENCE
// Extensible
type SearchspaceextV1700 struct {
	MonitoringslotperiodicityandoffsetV1710 *SearchspaceextV1700MonitoringslotperiodicityandoffsetV1710
	MonitoringslotswithinslotgroupR17       *SearchspaceextV1700MonitoringslotswithinslotgroupR17
	DurationR17                             *utils.INTEGER `lb:0,ub:20476`
	SearchspacetypeR17                      *SearchspaceextV1700SearchspacetypeR17
	SearchspacegroupidlistR17               *[]utils.INTEGER `lb:1,ub:3`
	SearchspacelinkingidR17                 *utils.INTEGER   `lb:0,ub:maxNrofSearchSpacesLinks1R17`
}
