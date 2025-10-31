package ies

import "rrc/utils"

// PDCCH-ServingCellConfig ::= SEQUENCE
// Extensible
type PdcchServingcellconfig struct {
	Slotformatindicator         *utils.Setuprelease[Slotformatindicator]
	AvailabilityindicatorR16    *utils.Setuprelease[AvailabilityindicatorR16] `ext`
	SearchspaceswitchtimerR16   *utils.INTEGER                                `lb:0,ub:80,ext`
	SearchspaceswitchtimerV1710 *utils.INTEGER                                `lb:0,ub:1280,ext`
}
