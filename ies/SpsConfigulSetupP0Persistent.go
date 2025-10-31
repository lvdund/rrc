package ies

import "rrc/utils"

// SPS-ConfigUL-setup-p0-Persistent ::= SEQUENCE
type SpsConfigulSetupP0Persistent struct {
	P0NominalpuschPersistent utils.INTEGER `lb:0,ub:24`
	P0UePuschPersistent      utils.INTEGER `lb:0,ub:7`
}
