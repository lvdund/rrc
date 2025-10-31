package ies

import "rrc/utils"

// IAB-ResourceConfig-r17 ::= SEQUENCE
// Extensible
type IabResourceconfigR17 struct {
	IabResourceconfigidR17       IabResourceconfigidR17
	SlotlistR17                  *[]utils.INTEGER `lb:1,ub:5120`
	PeriodicityslotlistR17       *IabResourceconfigR17PeriodicityslotlistR17
	SlotlistsubcarrierspacingR17 *Subcarrierspacing
}
