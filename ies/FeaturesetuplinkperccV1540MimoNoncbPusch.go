package ies

import "rrc/utils"

// FeatureSetUplinkPerCC-v1540-mimo-NonCB-PUSCH ::= SEQUENCE
type FeaturesetuplinkperccV1540MimoNoncbPusch struct {
	MaxnumbersrsResourceperset         utils.INTEGER `lb:0,ub:4`
	MaxnumbersimultaneoussrsResourcetx utils.INTEGER `lb:0,ub:4`
}
