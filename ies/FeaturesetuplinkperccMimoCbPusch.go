package ies

import "rrc/utils"

// FeatureSetUplinkPerCC-mimo-CB-PUSCH ::= SEQUENCE
type FeaturesetuplinkperccMimoCbPusch struct {
	MaxnumbermimoLayerscbPusch *MimoLayersul
	MaxnumbersrsResourceperset utils.INTEGER `lb:0,ub:2`
}
