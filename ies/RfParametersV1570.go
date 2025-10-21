package ies

import "rrc/utils"

// RF-Parameters-v1570 ::= SEQUENCE
type RfParametersV1570 struct {
	Dl1024qamScalingfactorR15       utils.ENUMERATED
	Dl1024qamTotalweightedlayersR15 utils.INTEGER
}
