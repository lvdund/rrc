package ies

import "rrc/utils"

// RF-Parameters-v1570 ::= SEQUENCE
type RfParametersV1570 struct {
	Dl1024qamScalingfactorR15       RfParametersV1570Dl1024qamScalingfactorR15
	Dl1024qamTotalweightedlayersR15 utils.INTEGER `lb:0,ub:10`
}
