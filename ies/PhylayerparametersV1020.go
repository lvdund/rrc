package ies

import "rrc/utils"

// PhyLayerParameters-v1020 ::= SEQUENCE
type PhylayerparametersV1020 struct {
	TwoantennaportsforpucchR10       *utils.ENUMERATED
	Tm9With8txFddR10                 *utils.ENUMERATED
	PmiDisablingR10                  *utils.ENUMERATED
	CrosscarrierschedulingR10        *utils.ENUMERATED
	SimultaneouspucchPuschR10        *utils.ENUMERATED
	MulticlusterpuschWithinccR10     *utils.ENUMERATED
	NoncontiguousulRaWithinccListR10 *NoncontiguousulRaWithinccListR10
}
