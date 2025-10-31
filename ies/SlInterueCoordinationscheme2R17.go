package ies

import "rrc/utils"

// SL-InterUE-CoordinationScheme2-r17 ::= SEQUENCE
// Extensible
type SlInterueCoordinationscheme2R17 struct {
	SlIucScheme2R17                 *SlInterueCoordinationscheme2R17SlIucScheme2R17
	SlRbSetpsfchR17                 *utils.BITSTRING `lb:10,ub:275`
	SlTypeueAR17                    *SlInterueCoordinationscheme2R17SlTypeueAR17
	SlPsfchOccasionR17              *utils.INTEGER `lb:0,ub:1`
	SlSlotlevelresourceexclusionR17 *SlInterueCoordinationscheme2R17SlSlotlevelresourceexclusionR17
	SlOptionforcondition2A1R17      *utils.INTEGER `lb:0,ub:1`
	SlIndicationueBR17              *SlInterueCoordinationscheme2R17SlIndicationueBR17
	SlDeltarsrpThreshV1720          *utils.INTEGER `lb:0,ub:30,ext`
}
