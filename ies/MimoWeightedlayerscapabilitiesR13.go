package ies

import "rrc/utils"

// MIMO-WeightedLayersCapabilities-r13 ::= SEQUENCE
type MimoWeightedlayerscapabilitiesR13 struct {
	RelweighttwolayersR13   MimoWeightedlayerscapabilitiesR13RelweighttwolayersR13
	RelweightfourlayersR13  *MimoWeightedlayerscapabilitiesR13RelweightfourlayersR13
	RelweighteightlayersR13 *MimoWeightedlayerscapabilitiesR13RelweighteightlayersR13
	TotalweightedlayersR13  utils.INTEGER `lb:0,ub:128`
}
