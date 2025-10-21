package ies

import "rrc/utils"

// MIMO-WeightedLayersCapabilities-r13 ::= SEQUENCE
type MimoWeightedlayerscapabilitiesR13 struct {
	RelweighttwolayersR13   utils.ENUMERATED
	RelweightfourlayersR13  *utils.ENUMERATED
	RelweighteightlayersR13 *utils.ENUMERATED
	TotalweightedlayersR13  utils.INTEGER
}
