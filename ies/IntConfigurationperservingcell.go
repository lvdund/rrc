package ies

import "rrc/utils"

// INT-ConfigurationPerServingCell ::= SEQUENCE
type IntConfigurationperservingcell struct {
	Servingcellid Servcellindex
	Positionindci utils.INTEGER `lb:0,ub:maxINTDciPayloadsize1`
}
