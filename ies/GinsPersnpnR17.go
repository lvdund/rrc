package ies

import "rrc/utils"

// GINs-PerSNPN-r17 ::= SEQUENCE
type GinsPersnpnR17 struct {
	SupportedginsR17 *utils.BITSTRING `lb:1,ub:maxGINR17`
}
