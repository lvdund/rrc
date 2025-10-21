package ies

import "rrc/utils"

// PhyLayerParameters-NB-v1610 ::= SEQUENCE
type PhylayerparametersNbV1610 struct {
	NpdschMultitbR16             *utils.ENUMERATED
	NpdschMultitbInterleavingR16 *utils.ENUMERATED
	NpuschMultitbR16             *utils.ENUMERATED
	NpuschMultitbInterleavingR16 *utils.ENUMERATED
	MultitbHarqAckbundlingR16    *utils.ENUMERATED
	SlotsymbolresourceresvdlR16  *utils.ENUMERATED
	SlotsymbolresourceresvulR16  *utils.ENUMERATED
	SubframeresourceresvdlR16    *utils.ENUMERATED
	SubframeresourceresvulR16    *utils.ENUMERATED
}
