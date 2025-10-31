package ies

import "rrc/utils"

// SPS-ConfigUL-STTI-r15-setup-p0-Persistent-r15 ::= SEQUENCE
type SpsConfigulSttiR15SetupP0PersistentR15 struct {
	P0NominalspuschPersistentR15 utils.INTEGER `lb:0,ub:24`
	P0UeSpuschPersistentR15      utils.INTEGER `lb:0,ub:7`
}
