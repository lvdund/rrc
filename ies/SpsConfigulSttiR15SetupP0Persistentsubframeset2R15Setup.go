package ies

import "rrc/utils"

// SPS-ConfigUL-STTI-r15-setup-p0-PersistentSubframeSet2-r15-setup ::= SEQUENCE
type SpsConfigulSttiR15SetupP0Persistentsubframeset2R15Setup struct {
	P0NominalspuschPersistentsubframeset2R15 utils.INTEGER `lb:0,ub:24`
	P0UeSpuschPersistentsubframeset2R15      utils.INTEGER `lb:0,ub:7`
}
